import React, {Fragment} from 'react';
import {Box, Button, Paper, Space, Text, Title} from "@mantine/core";
import {sendScanImage} from "./api.js";
import {ScanJobsTable} from "../../../components/ScanJobsTable.jsx";
import {Link} from "react-router-dom";
import Filters from "../../../components/Filters.jsx";

const ScannerScanImageView = () => {

    const runScan = () => {
        sendScanImage("test")
            .then(r => console.log(r))
            .catch(e => console.log(e));
    }

    const breadcrumbs = [
        {title: 'Scans', href: '/scans'},
        {title: 'Jobs', href: '/scans'},
    ].map((item, index) => (
        <Text component={Link} variant="link" to={item.href} key={index}>
            {item.title}
        </Text>
    ));

    return (
        <Fragment>
            <Box p="lg">
                <Paper bg="none">
                    <Title fw="lighter" size="xx-large" mb="xl">Scans</Title>
                </Paper>

                <Paper padding="md" radius={0} style={{padding: "30px"}}>
                    <Button onClick={runScan}>
                        Run Scan
                    </Button>
                    {/*<Filters onChange={onFilterChange}/>*/}
                </Paper>

                <Space h="xl"/>

                <ScanJobsTable />
            </Box>
        </Fragment>
    );
}

export default ScannerScanImageView;
