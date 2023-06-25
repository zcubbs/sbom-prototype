import React, {Fragment} from 'react';
import {Box, Paper, Space, Text, Title} from "@mantine/core";
import {Link} from "react-router-dom";
import RunScanForm from "../features/scanner/scan-image/RunScanForm.jsx";
import {JobsTable} from "../features/scanner/list-jobs/JobsTable.jsx";

const ScanView = () => {
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
                    <RunScanForm />
                </Paper>

                <Space h="xl"/>

                <JobsTable />
            </Box>
        </Fragment>
    );
};

export default ScanView;
