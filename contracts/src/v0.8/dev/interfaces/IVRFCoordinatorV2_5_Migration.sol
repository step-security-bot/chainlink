// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

// Future versions of VRFCoordinatorV2_5 must implement IVRFCoordinatorV2_5_Migration
// to support migrations from previous versions
interface IVRFCoordinatorV2_5_Migration {
  /**
   * @notice called by older versions of coordinator for migration.
   * @notice only callable by older versions of coordinator
   * @notice supports transfer of native currency
   * @param encodedData - user data from older version of coordinator
   */
  function onMigration(bytes calldata encodedData) external payable;
}