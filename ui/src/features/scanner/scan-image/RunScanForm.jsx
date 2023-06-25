import React from 'react';
import {useForm} from '@mantine/form';
import {Button, Grid, TextInput} from "@mantine/core";
import {sendScanImage} from "../api.js";

const RunScanForm = () => {
    const form = useForm({
        initialValues: {
            image: '',
        },

        validate: {
            image: (value) => (value.length > 0 ? null : 'Provide a value'),
        },
    });

    const onsubmit = (values) => {
        sendScanImage(values.image)
            .then(r => console.log(r))
            .catch(e => console.log(e));
    }

    return (
        <form onSubmit={form.onSubmit((values) => onsubmit(values))}>
            <Grid>
                <Grid.Col span={4}>
                    <TextInput
                        label="Image"
                        placeholder="ubuntu:latest"
                        {...form.getInputProps('image')}
                    />
                </Grid.Col>
                <Grid.Col span={4}>
                    <Button type="submit" mt="25px">Run Scan</Button>
                </Grid.Col>
            </Grid>
        </form>
    );
};

export default RunScanForm;
