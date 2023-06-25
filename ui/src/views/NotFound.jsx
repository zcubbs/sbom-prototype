import React, {Fragment} from "react";
import {Box} from "@mantine/core";
import NotFoundSplash from "../components/NotFoundSplash.jsx";

const NotFound = () => {

    const stats = [
        {
            title: "Total Posts",
            stats: "100",
            description: "24% more than month last month"
        },
        {
            title: "Adoption Candidates",
            stats: "23",
            description: "50% more than last month"
        },
        {
            title: "Adoptions",
            stats: "47",
            description: "Cats adopted this year"
        },
    ]

    return (
        <Fragment>
            <Box p="lg">
                <NotFoundSplash />
            </Box>
        </Fragment>
    );
}

export default NotFound;
