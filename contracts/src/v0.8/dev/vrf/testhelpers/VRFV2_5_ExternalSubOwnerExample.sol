// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../../shared/interfaces/LinkTokenInterface.sol";
import "../../interfaces/IVRFCoordinatorV2_5.sol";
import "../VRFConsumerBaseV2_5.sol";

/// @notice This contract is used for testing only and should not be used for production.
contract VRFV2_5_ExternalSubOwnerExample is VRFConsumerBaseV2_5 {
  IVRFCoordinatorV2_5 COORDINATOR;
  LinkTokenInterface LINKTOKEN;

  uint256[] public s_randomWords;
  uint256 public s_requestId;
  address s_owner;

  constructor(address vrfCoordinator, address link) VRFConsumerBaseV2_5(vrfCoordinator) {
    COORDINATOR = IVRFCoordinatorV2_5(vrfCoordinator);
    LINKTOKEN = LinkTokenInterface(link);
    s_owner = msg.sender;
  }

  function fulfillRandomWords(uint256 requestId, uint256[] memory randomWords) internal override {
    require(requestId == s_requestId, "request ID is incorrect");
    s_randomWords = randomWords;
  }

  function requestRandomWords(
    uint256 subId,
    uint32 callbackGasLimit,
    uint16 requestConfirmations,
    uint32 numWords,
    bytes32 keyHash,
    bool nativePayment
  ) external onlyOwner {
    VRFV2_5_Client.RandomWordsRequest memory req = VRFV2_5_Client.RandomWordsRequest({
      keyHash: keyHash,
      subId: subId,
      requestConfirmations: requestConfirmations,
      callbackGasLimit: callbackGasLimit,
      numWords: numWords,
      extraArgs: VRFV2_5_Client._argsToBytes(VRFV2_5_Client.ExtraArgsV1({nativePayment: nativePayment}))
    });
    // Will revert if subscription is not funded.
    s_requestId = COORDINATOR.requestRandomWords(req);
  }
}
