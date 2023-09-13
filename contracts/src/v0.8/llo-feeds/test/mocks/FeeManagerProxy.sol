// SPDX-License-Identifier: MIT
pragma solidity 0.8.16;

import "../../dev/interfaces/IFeeManager.sol";

contract FeeManagerProxy {
  IFeeManager internal i_feeManager;

  function processFee(bytes calldata payload, bytes calldata feePayload) public payable {
    i_feeManager.processFee{value: msg.value}(payload, feePayload, msg.sender);
  }

  function processFeeBulk(bytes[] calldata payloads, bytes calldata feePayload) public payable {
    i_feeManager.processFeeBulk{value: msg.value}(payloads, feePayload, msg.sender);
  }

  function setFeeManager(IFeeManager feeManager) public {
    i_feeManager = feeManager;
  }
}
