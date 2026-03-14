<?php
// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
//
// Copyright The Stargate Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
namespace Stargate;

/**
 * The gPRC service to interact with a Stargate coordinator.
 */
class StargateClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * Executes a single CQL query.
     * @param \Stargate\Query $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall<\Stargate\Response>
     */
    public function ExecuteQuery(\Stargate\Query $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/stargate.Stargate/ExecuteQuery',
        $argument,
        ['\Stargate\Response', 'decode'],
        $metadata, $options);
    }

    /**
     * Executes a bi-directional streaming for CQL queries.
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function ExecuteQueryStream($metadata = [], $options = []) {
        return $this->_bidiRequest('/stargate.Stargate/ExecuteQueryStream',
        ['\Stargate\StreamingResponse','decode'],
        $metadata, $options);
    }

    /**
     * Executes a batch of CQL queries.
     * @param \Stargate\Batch $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall<\Stargate\Response>
     */
    public function ExecuteBatch(\Stargate\Batch $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/stargate.Stargate/ExecuteBatch',
        $argument,
        ['\Stargate\Response', 'decode'],
        $metadata, $options);
    }

    /**
     * Executes a bi-directional streaming for batches of CQL queries.
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function ExecuteBatchStream($metadata = [], $options = []) {
        return $this->_bidiRequest('/stargate.Stargate/ExecuteBatchStream',
        ['\Stargate\StreamingResponse','decode'],
        $metadata, $options);
    }

}
