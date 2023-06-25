import {Box, Navbar as MantineNavbar, NavLink} from '@mantine/core';
import {useEffect, useState} from "react";
import {IconActivity, IconAnalyze, IconFingerprint, IconGauge} from "@tabler/icons-react";
import {useNavigate, useLocation} from "react-router-dom";

const data = [
    {
        icon: IconGauge,
        color: 'orange',
        label: 'Scans',
        description: 'Sbom & Image scan Jobs',
        rightSection: <IconActivity size={18} stroke={1.5}/>,
        disableRightSectionRotation: true,
        to: '/scans',
    },
    {
        icon: IconFingerprint,
        color: 'red',
        label: 'SBOM Registry',
        description: 'Run scans from sboms',
        rightSection: <IconAnalyze size={18} stroke={1.5}/>,
        disableRightSectionRotation: true,
        to: '/registry',
    },
];

export default function Navbar({opened}) {
    const [active, setActive] = useState(0);
    const navigate = useNavigate();
    const {pathname} = useLocation()

    useEffect(() => {
        data.map((item, index) => {
            if (item.to === pathname) {
                setActive(index);
            }
        })
    }, [opened]);

    const items = data.map((item, index) => (
        <NavLink fw="400"
                 key={item.label}
                 active={index === active}
                 label={item.label}
                 description={item.description}
                 rightSection={item.rightSection}
                 icon={<item.icon size={16} stroke={1.5} color={item.color}/>}
                 onClick={() => {
                     setActive(index);
                     navigate(item.to);
                 }}
                 variant="light"
        />
    ));

    return (
        <MantineNavbar hiddenBreakpoint="sm" hidden={!opened} width={{sm: 200, lg: 300}}>
            <MantineNavbar.Section>
                <Box>
                    {items}
                </Box>
            </MantineNavbar.Section>
        </MantineNavbar>
    );
}
