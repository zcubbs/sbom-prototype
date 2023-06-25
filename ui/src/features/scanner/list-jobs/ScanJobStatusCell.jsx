import React from 'react';
import {Loader, Text} from "@mantine/core";

const ScanJobStatusCell = (value) => {
    if (value.getValue() === "pending") {
        return (
            <Loader color="orange" size="sm" variant="dots"/>
        );
    }

    if (value.getValue() === "running") {
        return (
            <Loader size="sm"/>
        );
    }

    return (
        <Text size="sm">?</Text>
    );
};

export default ScanJobStatusCell;
