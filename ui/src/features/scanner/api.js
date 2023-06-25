import axios from "axios";
import {API_URL} from "../../config/config.js";

export const sendScanImage = async (image) => {
    const response = await axios
        .post(API_URL + '/v1/scan/image', {
            image: image,
        });

    return response.data;
}

// export const fetchScans = async () => {
//     const response = await axios.get(API_URL + '/v1/scans');
//     console.log(response.data);
//     return response.data;
// }

export const fetchScans = async ({
                                     pageIndex,
                                     pageSize,
                                     givenImage,
                                     scoreGreaterThan
                                 }) => {
    const response = await fetch(API_URL + `/v1/scans`
        + `?page=${pageIndex}&page_size=${pageSize}&page_order=asc`,
        {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'GET',
            body: JSON.stringify(
                {
                    image: givenImage,
                    score_greater_than: scoreGreaterThan,
                }
            )
        });
    return response.json();
};
