import React from "react";
import {Flex, Loader, Text} from '@mantine/core';

const Loading = () => (
    <Flex
        mt="45vh"
        mih={50}
        gap="md"
        justify="center"
        align="center"
        direction="column"
        wrap="wrap"
    >
        <Loader size={50} position="center" mt="auto"/>
        <Text size="xl" weight={500} mt="auto">
            Loading...
        </Text>
    </Flex>
);

export default Loading;
