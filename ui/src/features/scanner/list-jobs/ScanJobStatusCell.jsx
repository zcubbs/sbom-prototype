import React from 'react';
import {Loader, Text, ThemeIcon} from "@mantine/core";
import {IconClockPause} from "@tabler/icons-react";

const ScanJobStatusCell = (value) => {
    if (value.getValue() === "pending") {
        return (
            <ThemeIcon size="lg" color="gray">
                <IconClockPause size="1.2rem"/>
            </ThemeIcon>
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
