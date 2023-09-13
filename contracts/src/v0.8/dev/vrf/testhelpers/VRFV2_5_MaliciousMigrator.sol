// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../../interfaces/IVRFMigratableConsumerV2_5.sol";
import "../../interfaces/IVRFCoordinatorV2_5.sol";
import "../libraries/VRFV2_5_Client.sol";

contract VRFV2_5_MaliciousMigrator is IVRFMigratableConsumerV2_5 {
  IVRFCoordinatorV2_5 s_vrfCoordinator;

  constructor(address _vrfCoordinator) {
    s_vrfCoordinator = IVRFCoordinatorV2_5(_vrfCoordinator);
  }

  /**
   * @inheritdoc IVRFMigratableConsumerV2_5
   */
  function setCoordinator(address /*_vrfCoordinator*/) public override {
    // try to re-enter, should revert
    // args don't really matter
    s_vrfCoordinator.requestRandomWords(
      VRFV2_5_Client.RandomWordsRequest({
        keyHash: bytes32(0),
        subId: 0,
        requestConfirmations: 0,
        callbackGasLimit: 0,
        numWords: 0,
        extraArgs: ""
      })
    );
  }
}
