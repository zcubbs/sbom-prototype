import React from 'react';
import {API_URL} from "../../config/config.js";

const RegistryIndex = () => {
    const {isLoading, error, data} = useQuery({
        queryKey: ['repoData'],
        queryFn: () =>
            fetch(API_URL + '/v1/scanner/list').then(
                (res) => res.json(),
            ),
    })

    if (isLoading) return 'Loading...'
    if (error) return 'An error has occurred: ' + error.message

    return (
        <div className="App">{data}</div>
    )
};

export default RegistryIndex;
