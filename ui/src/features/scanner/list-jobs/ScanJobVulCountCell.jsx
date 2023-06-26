import React from 'react';
import {Badge} from "@mantine/core";

const ScanJobVulCountCell = (value) => {
    if (value.getValue() === undefined) {
        return (
            <></>
        );
    }
    return (
        <Badge>{value.getValue()}</Badge>
    );
};

export default ScanJobVulCountCell;
