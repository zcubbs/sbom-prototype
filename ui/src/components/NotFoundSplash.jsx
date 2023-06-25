import {createStyles, Image, Container, Title, Text, Button, SimpleGrid} from '@mantine/core';
import image from '../assets/404.svg';
import * as React from "react";

const useStyles = createStyles((theme) => ({
    root: {
        paddingTop: 80,
        paddingBottom: 80,
    },

    title: {
        fontWeight: 900,
        fontSize: 34,
        marginBottom: theme.spacing.md,
        fontFamily: `Greycliff CF, ${theme.fontFamily}`,

        [theme.fn.smallerThan('sm')]: {
            fontSize: 32,
        },
    },

    control: {
        [theme.fn.smallerThan('sm')]: {
            width: '100%',
        },
    },

    mobileImage: {
        [theme.fn.largerThan('sm')]: {
            display: 'none',
        },
    },

    desktopImage: {
        [theme.fn.smallerThan('sm')]: {
            display: 'none',
        },
    },
}));

export default function NotFoundSplash() {
    const {classes, theme} = useStyles();

    const goHome = () => {
        window.location.href = '/';
    }

    return (
        <Container className={classes.root}>
            <SimpleGrid spacing={80} cols={2} breakpoints={[{maxWidth: 'sm', cols: 1, spacing: 40}]}>
                <Image src={image} className={classes.mobileImage}/>
                <div>
                    {theme.colorScheme === 'dark' ?
                        <Title style={{color: '#fff' }} className={classes.title}>Something is not right...</Title>
                        :
                        <Title className={classes.title}>Something is not right...</Title>
                    }

                    <Text color="dimmed" size="lg">
                        Page you are trying to open does not exist. You may have mistyped the address, or the
                        page has been moved to another URL. If you think this is an error contact support.
                    </Text>
                    <Button variant="outline" size="md" mt="xl" className={classes.control} onClick={goHome}>
                        Get back to home page
                    </Button>
                </div>
                <Image src={image} className={classes.image}/>
            </SimpleGrid>
        </Container>
    );
}
