import {AppShell, useMantineTheme} from '@mantine/core';
import Header from "./Header.jsx";
import Navbar from "./Navbar.jsx";
import * as React from "react";
import {useState} from "react";
import {Route, Routes} from "react-router-dom";
import NotFound from "../views/NotFound.jsx";
import ScanView from "../views/ScanView.jsx";

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
                <Route path="/" element={<ScanView/>}/>
                <Route path="/scans" element={<ScanView/>}/>
                <Route path="/registry" element={<NotFound/>}/>
            </Routes>

        </AppShell>
    );
}
