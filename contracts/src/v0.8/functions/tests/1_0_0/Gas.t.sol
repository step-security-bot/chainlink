// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {BaseTest} from "./BaseTest.t.sol";
import {FunctionsRouter} from "../../dev/1_0_0/FunctionsRouter.sol";
import {FunctionsSubscriptions} from "../../dev/1_0_0/FunctionsSubscriptions.sol";
import {FunctionsRequest} from "../../dev/1_0_0/libraries/FunctionsRequest.sol";
import {FunctionsResponse} from "../../dev/1_0_0/libraries/FunctionsResponse.sol";
import {FunctionsClientTestHelper} from "./testhelpers/FunctionsClientTestHelper.sol";

import {FunctionsRoutesSetup, FunctionsOwnerAcceptTermsOfServiceSetup, FunctionsSubscriptionSetup, FunctionsClientRequestSetup} from "./Setup.t.sol";

import "forge-std/Vm.sol";
import "forge-std/console.sol";

/// @notice #acceptTermsOfService
contract Gas_AcceptTermsOfService is FunctionsRoutesSetup {
  bytes32 s_sigR;
  bytes32 s_sigS;
  uint8 s_sigV;

  function setUp() public virtual override {
    FunctionsRoutesSetup.setUp();

    bytes32 message = s_termsOfServiceAllowList.getMessage(OWNER_ADDRESS, OWNER_ADDRESS);
    bytes32 prefixedMessage = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", message));
    (s_sigV, s_sigR, s_sigS) = vm.sign(TOS_SIGNER_PRIVATE_KEY, prefixedMessage);
  }

  function test_AcceptTermsOfService_Gas() public {
    s_termsOfServiceAllowList.acceptTermsOfService(OWNER_ADDRESS, OWNER_ADDRESS, s_sigR, s_sigS, s_sigV);
  }
}

/// @notice #createSubscription
contract Gas_CreateSubscription is FunctionsOwnerAcceptTermsOfServiceSetup {
  function test_CreateSubscription_Gas() public {
    s_functionsRouter.createSubscription();
  }
}

/// @notice #addConsumer
contract Gas_AddConsumer is FunctionsSubscriptionSetup {
  // use garbage address
  address consumerAddress = address(s_functionsCoordinator);

  function test_AddConsumer_Gas() public {
    s_functionsRouter.addConsumer(s_subscriptionId, consumerAddress);
  }
}

/// @notice #fundSubscription
contract Gas_FundSubscription is FunctionsSubscriptionSetup {
  address routerAddress = address(s_functionsRouter);
  uint96 s_subscriptionFunding = 10 * JUELS_PER_LINK; // 10 LINK
  bytes data = abi.encode(s_subscriptionId);

  function test_FundSubscription_Gas() public {
    s_linkToken.transferAndCall(routerAddress, s_subscriptionFunding, data);
  }
}

/// @notice #sendRequest
contract Gas_SendRequest is FunctionsSubscriptionSetup {
  bytes s_minimalRequestData;
  bytes s_maximalRequestData;

  function _makeStringOfBytesSize(uint16 bytesSize) internal pure returns (string memory) {
    return vm.toString(new bytes((bytesSize - 2) / 2));
  }

  function setUp() public virtual override {
    FunctionsSubscriptionSetup.setUp();

    {
      // Create minimum viable request data
      FunctionsRequest.Request memory minimalRequest;
      string memory minimalSourceCode = "return Functions.encodeString('hello world');";
      FunctionsRequest.initializeRequest(
        minimalRequest,
        FunctionsRequest.Location.Inline,
        FunctionsRequest.CodeLanguage.JavaScript,
        minimalSourceCode
      );
      s_minimalRequestData = FunctionsRequest.encodeCBOR(minimalRequest);
    }

    {
      // Create maximum viable request data - 30 KB encoded data
      FunctionsRequest.Request memory maxmimalRequest;

      // Create maximum viable request data - 30 KB encoded data
      string memory maximalSourceCode = _makeStringOfBytesSize(29_898); // CBOR size without source code is 102 bytes
      FunctionsRequest.initializeRequest(
        maxmimalRequest,
        FunctionsRequest.Location.Inline,
        FunctionsRequest.CodeLanguage.JavaScript,
        maximalSourceCode
      );
      s_maximalRequestData = FunctionsRequest.encodeCBOR(maxmimalRequest);
      assertEq(s_maximalRequestData.length, 30_000);
    }
  }

  /// @dev The order of these test cases matters as the first test will consume more gas by writing over default values
  function test_SendRequest_MaximumGas() public {
    s_functionsClient.sendRequestBytes(s_maximalRequestData, s_subscriptionId, 300_000, s_donId);
  }

  function test_SendRequest_MinimumGas() public {
    s_functionsClient.sendRequestBytes(s_minimalRequestData, s_subscriptionId, 5_500, s_donId);
  }
}

/// @notice #fulfillRequest
contract FunctionsClient_FulfillRequest is FunctionsClientRequestSetup {
  struct Report {
    bytes32[] rs;
    bytes32[] ss;
    bytes32 vs;
    bytes report;
    bytes32[3] reportContext;
  }

  mapping(uint256 reportNumber => Report) s_reports;

  FunctionsClientTestHelper s_functionsClientWithMaximumReturnData;

  function _makeStringOfBytesSize(uint16 bytesSize) internal pure returns (string memory) {
    return vm.toString(new bytes((bytesSize - 2) / 2));
  }

  function setUp() public virtual override {
    FunctionsSubscriptionSetup.setUp();

    {
      // Deploy consumer that has large revert return data
      s_functionsClientWithMaximumReturnData = new FunctionsClientTestHelper(address(s_functionsRouter));
      s_functionsClientWithMaximumReturnData.setRevertFulfillRequest(true);
      string memory revertMessage = _makeStringOfBytesSize(30_000); // 30kb - FunctionsRouter cuts off response at MAX_CALLBACK_RETURN_BYTES = 4 + 4 * 32 = 132bytes, go well above that
      s_functionsClientWithMaximumReturnData.setRevertFulfillRequestMessage(revertMessage);
      s_functionsRouter.addConsumer(s_subscriptionId, address(s_functionsClientWithMaximumReturnData));
    }

    // Set up maximum gas test
    {
      // Send request #2 for maximum gas test
      uint8 requestNumber = 2;

      bytes memory secrets = new bytes(0);
      string[] memory args = new string[](0);
      bytes[] memory bytesArgs = new bytes[](0);
      uint32 callbackGasLimit = 300_000;

      // Create maximum viable request data - 30 KB encoded data
      string memory maximalSourceCode = _makeStringOfBytesSize(29_898); // CBOR size without source code is 102 bytes

      _sendAndStoreRequest(
        requestNumber,
        maximalSourceCode,
        secrets,
        args,
        bytesArgs,
        callbackGasLimit,
        address(s_functionsClientWithMaximumReturnData)
      );

      // Build the report transmission data
      uint256[] memory requestNumberKeys = new uint256[](1);
      requestNumberKeys[0] = requestNumber;
      string[] memory results = new string[](1);
      // Build a 256 byte response size
      results[0] = _makeStringOfBytesSize(256);
      bytes[] memory errors = new bytes[](1);
      errors[0] = new bytes(0); // No error

      (bytes memory report, bytes32[3] memory reportContext) = _buildReport(requestNumberKeys, results, errors);

      uint256[] memory signerPrivateKeys = new uint256[](3);
      signerPrivateKeys[0] = NOP_SIGNER_PRIVATE_KEY_1;
      signerPrivateKeys[1] = NOP_SIGNER_PRIVATE_KEY_2;
      signerPrivateKeys[2] = NOP_SIGNER_PRIVATE_KEY_3;

      (bytes32[] memory rawRs, bytes32[] memory rawSs, bytes32 rawVs) = _signReport(
        report,
        reportContext,
        signerPrivateKeys
      );

      // Store the report data
      s_reports[1] = Report({rs: rawRs, ss: rawSs, vs: rawVs, report: report, reportContext: reportContext});
    }

    // Set up minimum gas test
    {
      // Send requests minimum gas test
      uint8 requestsToSend = 1;
      uint8 requestNumberOffset = 3; // the setup already has request #1 sent, and the previous test case uses request #2, start from request #3

      string memory sourceCode = "return Functions.encodeString('hello world');";
      bytes memory secrets = new bytes(0);
      string[] memory args = new string[](0);
      bytes[] memory bytesArgs = new bytes[](0);
      uint32 callbackGasLimit = 10_000;

      for (uint256 i = 0; i < requestsToSend; ++i) {
        _sendAndStoreRequest(i + requestNumberOffset, sourceCode, secrets, args, bytesArgs, callbackGasLimit);
      }

      // Build the report transmission data
      uint256[] memory requestNumberKeys = new uint256[](requestsToSend);
      string[] memory results = new string[](requestsToSend);
      bytes[] memory errors = new bytes[](requestsToSend);
      for (uint256 i = 0; i < requestsToSend; ++i) {
        requestNumberKeys[i] = i + requestNumberOffset;
        results[i] = "hello world";
        errors[i] = new bytes(0); // no error
      }

      (bytes memory report, bytes32[3] memory reportContext) = _buildReport(requestNumberKeys, results, errors);

      uint256[] memory signerPrivateKeys = new uint256[](3);
      signerPrivateKeys[0] = NOP_SIGNER_PRIVATE_KEY_1;
      signerPrivateKeys[1] = NOP_SIGNER_PRIVATE_KEY_2;
      signerPrivateKeys[2] = NOP_SIGNER_PRIVATE_KEY_3;

      (bytes32[] memory rawRs, bytes32[] memory rawSs, bytes32 rawVs) = _signReport(
        report,
        reportContext,
        signerPrivateKeys
      );

      // Store the report data
      s_reports[2] = Report({rs: rawRs, ss: rawSs, vs: rawVs, report: report, reportContext: reportContext});
    }

    vm.stopPrank();
    vm.startPrank(NOP_TRANSMITTER_ADDRESS_1);
  }

  /// @dev The order of these test cases matters as the first test will consume more gas by writing over default values
  function test_FulfillRequest_MaximumGas() public {
    // 1 fulfillment in the report, single request takes on all report validation cost
    // maximum request
    // maximum NOPs
    // maximum return data
    // first storage write to change default values
    s_functionsCoordinator.transmit(
      s_reports[1].reportContext,
      s_reports[1].report,
      s_reports[1].rs,
      s_reports[1].ss,
      s_reports[1].vs
    );
  }

  function test_FulfillRequest_MinimumGas() public {
    // max fulfillments in the report, cost of validation split between all
    // minimal request
    // minimum NOPs
    // no return data
    // not storage writing default values
    s_functionsCoordinator.transmit(
      s_reports[2].reportContext,
      s_reports[2].report,
      s_reports[2].rs,
      s_reports[2].ss,
      s_reports[2].vs
    );
  }
}
