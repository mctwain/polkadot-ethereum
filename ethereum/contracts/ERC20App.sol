// SPDX-License-Identifier: MIT
pragma solidity >=0.6.2;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./Decoder.sol";
import "./Application.sol";
import "./SendChannel.sol";

contract ERC20App {
    using SafeMath for uint256;
    using Decoder for bytes;

    uint64 constant PAYLOAD_LENGTH = 104;
    string constant TARGET_APPLICATION_ID = "erc20-app";

    address public bridge;
    mapping(address => uint256) public totalTokens;
    address public basicSendChannelAddress;
    address public incentivizedSendChannelAddress;

    event Locked(
        address _sender,
        bytes32 _recipient,
        address _token,
        uint256 _amount
    );
    event Unlocked(
        bytes32 _polkadotSender,
        address _recipient,
        address _token,
        uint256 _amount
    );

    struct ERC20LockedPayload {
        address _sender;
        bytes32 _recipient;
        address _token;
        uint256 _amount;
    }

    constructor(
        address _basicSendChannelAddress,
        address _incentivizedSendChannelAddress
    ) public {
        basicSendChannelAddress = _basicSendChannelAddress;
        incentivizedSendChannelAddress = _incentivizedSendChannelAddress;
    }

    function register(address _bridge) public {
        require(bridge == address(0), "Bridge has already been registered");
        bridge = _bridge;
    }

    function sendERC20(
        bytes32 _recipient,
        address _tokenAddr,
        uint256 _amount,
        bool incentivized
    ) public {
        require(
            IERC20(_tokenAddr).transferFrom(msg.sender, address(this), _amount),
            "Contract token allowances insufficient to complete this lock request"
        );

        // Increment locked ERC20 token counter by this amount
        totalTokens[_tokenAddr] = totalTokens[_tokenAddr].add(_amount);

        emit Locked(msg.sender, _recipient, _tokenAddr, _amount);

        ERC20LockedPayload memory payload =
            ERC20LockedPayload(msg.sender, _recipient, _tokenAddr, _amount);
        SendChannel sendChannel;
        if (incentivized) {
            sendChannel = SendChannel(incentivizedSendChannelAddress);
        } else {
            sendChannel = SendChannel(basicSendChannelAddress);
        }
        sendChannel.send(TARGET_APPLICATION_ID, abi.encode(payload));
    }

    function sendTokens(
        bytes32 _polkadotSender,
        address _recipient,
        address _token,
        uint256 _amount
    ) public {
        require(msg.sender == bridge);
        require(_amount > 0, "Must unlock a positive amount");
        require(
            _amount <= totalTokens[_token],
            "ERC20 token balances insufficient to fulfill the unlock request"
        );

        totalTokens[_token] = totalTokens[_token].sub(_amount);
        require(
            IERC20(_token).transfer(_recipient, _amount),
            "ERC20 token transfer failed"
        );
        emit Unlocked(_polkadotSender, _recipient, _token, _amount);
    }
}
