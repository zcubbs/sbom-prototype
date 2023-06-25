import {AppShell, useMantineTheme} from '@mantine/core';
import Header from "./Header.jsx";
import Navbar from "./Navbar.jsx";
import * as React from "react";
import {useState} from "react";
import {Route, Routes} from "react-router-dom";
import NotFound from "../views/NotFound.jsx";
import Footer from "./Footer.jsx";
import ScannerScanImageView from "../features/scanner/scan-image/ScannerScanImageView.jsx";

export default function AppShellComponent() {
    const theme = useMantineTheme();
    const [opened, setOpened] = useState(false);

    return (
        <AppShell
            styles={{
                main: {
                    background: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.colors.gray[0],
                },
            }}
            navbarOffsetBreakpoint="sm"
            asideOffsetBreakpoint="sm"
            navbar={
                <Navbar opened={opened} />
            }
            // aside={
            //     <Sidebar/>
            // }
            // footer={
            //     <Footer/>
            // }
            header={
                <Header opened={opened} setOpened={setOpened} />
            }
        >
            <Routes>
                <Route path="/" element={<ScannerScanImageView/>}/>
                <Route path="/scans" element={<ScannerScanImageView/>}/>
                <Route path="/registry" element={<NotFound/>}/>
            </Routes>

        </AppShell>
    );
}
