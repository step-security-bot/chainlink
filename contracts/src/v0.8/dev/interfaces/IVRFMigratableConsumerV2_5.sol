// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice The IVRFMigratableConsumerV2_5 interface defines the
/// @notice method required to be implemented by all V2_5 consumers.
/// @dev This interface is designed to be used in VRFConsumerBaseV2_5.
interface IVRFMigratableConsumerV2_5 {
  /// @notice Sets the VRF Coordinator address
  /// @notice This method is should only be callable by the coordinator or contract owner
  function setCoordinator(address vrfCoordinator) external;
}
