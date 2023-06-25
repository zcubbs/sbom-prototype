import React, {Fragment} from 'react';
import {Box, Button} from "@mantine/core";
import {useQuery} from "@tanstack/react-query";
import {fetchScans, sendScanImage} from "./api.js";

const ScannerScanImageView = () => {
    const {isLoading, error, data} = useQuery({
        queryKey: ['repoData'],
        queryFn: () => fetchScans(),
    })

    if (isLoading) return 'Loading...'
    if (error) return 'An error has occurred: ' + error.message

    const runScan = () => {
        sendScanImage("test")
            .then(r => console.log(r))
            .catch(e => console.log(e));
    }

    return (
        <Fragment>
            <Box p="lg">
                <div className="App">
                    <p>{data.uuid}</p>
                    <p>{data.image}</p>
                </div>
                <Button onClick={runScan}>
                    Run Scan
                </Button>
            </Box>
        </Fragment>
    );
}

export default ScannerScanImageView;
