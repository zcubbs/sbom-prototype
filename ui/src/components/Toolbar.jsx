import {
    ActionIcon,
    Box,
    Burger,
    Center, createStyles,
    Group,
    MediaQuery,
    SegmentedControl, Text,
    useMantineColorScheme
} from '@mantine/core';
import {IconBolt, IconBrandGithubFilled, IconMoon, IconMoonStars, IconSun} from '@tabler/icons-react';
import * as React from "react";

export function Toolbar() {
    const {classes, cx, theme} = useStyles();
    const {colorScheme, toggleColorScheme} = useMantineColorScheme();

    return (
        <Box
            sx={(theme) => ({
                paddingLeft: theme.spacing.md,
                paddingRight: theme.spacing.md,
                paddingTop: theme.spacing.md,
            })}
        >
            <Group position="apart">
                {theme.colorScheme === 'dark' ?
                    <Text style={{color: '#fff' }}>
                        <IconBrandGithubFilled size={15} style={{marginRight: '5px'}}/>
                        <a style={{color: '#fff', marginLeft: '5px'}} href="https://github.com/zcubbs">github/zcubbs</a>
                    </Text>
                    :
                    <Text>
                        <IconBrandGithubFilled size={15} style={{marginRight: '5px'}}/>
                        <a style={{marginLeft: '5px'}}  href="https://github.com/zcubbs">github/zcubbs</a>
                    </Text>
                }

                <SegmentedControl
                    value={colorScheme}
                    onChange={(value) => toggleColorScheme(value)}
                    className={cx(classes.bigThemeToggle)}
                    data={[
                        {
                            value: 'light',
                            label: (
                                <Center>
                                    <IconSun size={16} stroke={1.5} />
                                    <Box ml={10}>Light</Box>
                                </Center>
                            ),
                        },
                        {
                            value: 'dark',
                            label: (
                                <Center>
                                    <IconMoon size={16} stroke={1.5} />
                                    <Box ml={10}>Dark</Box>
                                </Center>
                            ),
                        },
                    ]}
                />

                <MediaQuery largerThan="md" styles={{display: 'none'}}>
                    <ActionIcon variant="default" onClick={() => toggleColorScheme()} size={30}>
                        {colorScheme === 'dark' ? <IconSun size={16}/> : <IconMoonStars size={16}/>}
                    </ActionIcon>
                </MediaQuery>
            </Group>
        </Box>
    );
}

const useStyles = createStyles((theme) => ({
    bigThemeToggle: {
        [theme.fn.smallerThan('md')]: {
            display: 'none',
        },
    },
}));
