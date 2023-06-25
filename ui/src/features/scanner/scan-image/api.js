import axios from "axios";
import {API_URL} from "../../../config/config.js";

export const sendScanImage = async (image) => {
    const response = await axios
        .post(API_URL + '/v1/scan/image', {
            image: image,
        });

    return response.data;
}

export const fetchScans = async () => {
    const response = await axios.get(API_URL + '/v1/scans');
    return response.data;
}
