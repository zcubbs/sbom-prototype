import {createStyles, Flex, Text} from '@mantine/core';
import {IconBolt} from '@tabler/icons-react';

const useStyles = createStyles((theme) => ({
    footer: {
        borderTop: `1px solid ${
            theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
    },

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

    links: {
        [theme.fn.smallerThan('xs')]: {
            marginTop: theme.spacing.md,
        },
    },
}));

export default function Footer() {
    const { classes, theme } = useStyles();

    return (
        <div className={classes.footer}>
            <Flex className={classes.inner}
                  mih={50}
                  gap="md"
                  justify="center"
                  align="center"
                  direction="row"
                  wrap="wrap"
            >
                {theme.colorScheme === 'dark' ?
                    <Text style={{color: '#fff' }}>
                        <IconBolt size={15} style={{marginRight: '5px'}}/>
                        <a style={{color: '#fff', marginLeft: '5px'}} href="https://github.com/zcubbs">github/zcubbs</a>
                        - MIT {new Date().getFullYear()}
                    </Text>
                    :
                    <Text>
                        <IconBolt size={15} style={{marginRight: '5px'}}/>
                        <a style={{marginLeft: '5px'}}  href="https://github.com/zcubbs">github/zcubbs</a>
                        - MIT {new Date().getFullYear()}
                    </Text>
                }
            </Flex>
        </div>
    );
}
