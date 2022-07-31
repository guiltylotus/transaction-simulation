
// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;


contract IncreaseEvents {

    event NewID(
        uint256 date,
        uint256 id
    );
    event OldID(
        uint256 date,
        uint256 id
    );
    uint256 id;

    function increase() external  {
        emit OldID(block.timestamp, id);
        id++;
        emit NewID(block.timestamp, id);
    }

    function get() public view returns (uint256) {
        return id;
    }
}