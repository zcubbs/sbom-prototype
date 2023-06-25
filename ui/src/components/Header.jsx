import * as React from 'react';
import {Burger, Header as MantineHeader, MediaQuery, useMantineTheme} from '@mantine/core';
import {Brand} from "./Brand";
import {Toolbar} from "./Toolbar.jsx";

function Header({opened, setOpened}) {
    const theme = useMantineTheme();

    return (
        <MantineHeader height={70}>
            {/* Handle other responsive styles with MediaQuery component or createStyles function */}
            <div style={{
                display: 'flex',
                justifyContent: "space-between",
            }}>
                <MediaQuery largerThan="sm" styles={{display: 'none'}}>
                    <Burger
                        opened={opened}
                        onClick={() => setOpened((o) => !o)}
                        size="sm"
                        color={theme.colors.gray[6]}
                        mr="xl"
                        style={{
                            paddingLeft: '22px',
                            paddingRight: '16px',
                            paddingTop: '35px',
                        }}
                    />
                </MediaQuery>

                <Brand/>
                <Toolbar/>
            </div>
        </MantineHeader>
    );
}

export default Header;
