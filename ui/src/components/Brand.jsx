import {Box, createStyles, Group, Text, useMantineColorScheme} from '@mantine/core';
import {Logo} from './Logo';
import * as React from "react";

const useStyles = createStyles((theme) => ({
    inner: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        paddingTop: theme.spacing.xl,
        paddingBottom: theme.spacing.xl,

        [theme.fn.smallerThan('xs')]: {
            flexDirection: 'column',
        },
    },
}));

export function Brand() {
    const { theme } = useStyles();
    const {colorScheme} = useMantineColorScheme();

    return (
        <Box
            sx={(theme) => ({
                paddingLeft: theme.spacing.md,
                paddingRight: theme.spacing.md,
                paddingTop: theme.spacing.md,
            })}
        >
            <Group position="apart">
                <Logo colorScheme={colorScheme}/>
                {theme.colorScheme === 'dark' ?
                    <Text style={{color: '#fff' }}>
                        SBOM Prototype
                    </Text>
                    :
                    <Text>
                        SBOM Prototype
                    </Text>
                }
            </Group>
        </Box>
    );
}
