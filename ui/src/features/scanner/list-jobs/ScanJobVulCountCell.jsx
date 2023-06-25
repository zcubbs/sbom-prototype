import React from 'react';
import {Badge, Loader, Text, ThemeIcon} from "@mantine/core";
import {IconClockPause} from "@tabler/icons-react";

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
