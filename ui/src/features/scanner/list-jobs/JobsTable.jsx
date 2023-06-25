import {ScrollArea, Space} from "@mantine/core";
import {useNavigate} from "react-router-dom";
import {useQuery} from "@tanstack/react-query";
import {createColumnHelper, getCoreRowModel, useReactTable} from '@tanstack/react-table'
import {useMemo, useState} from "react";
import {API_URL} from "../../../config/config.js";
import PaginationWrapper from "../../../components/PaginationWrapper.jsx";
import TableWrapper from "../../../components/TableWrapper.jsx";
import ScanScoreCell from "./ScanScoreCell.jsx";
import ScanJobStatusCell from "./ScanJobStatusCell.jsx";

export const JobsTable = () => {
    const [scrolled, setScrolled] = useState(false);

    const navigate = useNavigate();

    const columnHelper = createColumnHelper();

    const [{pageIndex, pageSize}, setPagination] = useState({
        pageIndex: 0,
        pageSize: 10,
    });

    const [{
        scanId,
        image,
        scoreGreaterThan,
    }, setFilters] = useState({
        scanId: "",
        image: "",
        scoreGreaterThan: 0,
    })

    const [count, setCount] = useState(0);

    const fetchDataOptions = {
        pageIndex,
        pageSize,
        scanId,
        image,
        scoreGreaterThan,
    };

    const fetchScans = async ({
                                  pageIndex,
                                  pageSize,
                                  scanId,
                                  image,
                                  scoreGreaterThan
                              }) => {
        const response = await fetch(API_URL + `/v1/scans`
            + `?page=${pageIndex}&page_size=${pageSize}&page_order=asc`,
            {
                headers: {
                    'Content-Type': 'application/json',
                },
                method: 'GET',
            });
        return response.json();
    };

    const dataQuery = useQuery(
        ['data', fetchDataOptions],
        () => fetchScans(fetchDataOptions),
        {
            keepPreviousData: true,
            onSettled: (data) => {
                console.log(data);
                setCount(data.pagination.count);
            }
        }
    );

    const defaultData = useMemo(() => [], [])

    const pagination = useMemo(
        () => ({
            pageIndex,
            pageSize,
        }),
        [pageIndex, pageSize]
    )

    const filters = useMemo(() => ({
        scanId,
        image,
        scoreGreaterThan,
    }), [scanId, image, scoreGreaterThan])

    const columns = useMemo(() => [
        {
            header: "Status",
            accessorKey: "status",
            id: "status",
            cell: ScanJobStatusCell,
            size: 40,
        },
        {
            header: "Job Id",
            accessorKey: "uuid",
            id: "id",
            size: 350,
        },
        {
            header: "Risk Score",
            accessorKey: "risk_score",
            id: "risk_score",
            cell: ScanScoreCell,
            size: 80,
        },
        {
            header: "Vulnerabilities",
            accessorKey: "vulnerabilityCount",
            id: "vulnerability_count",
            size: 100,
        },
        {
            header: "Image",
            accessorKey: "image",
            id: "image",
        },
    ]);

    const tableHooks = (hooks) => {
        hooks.visibleColumns.push((columns) => [
            ...columns,
            {
                id: 'actions',
                header: 'actions',
                cell: ({row}) => (
                    <button onClick={() => goToScanReport(row.original)}>View</button>
                ),
            },
        ])
    }

    const table = useReactTable(
        {
            data: dataQuery.data?.scans ?? defaultData,
            columns,
            pageCount: dataQuery.data?.pagination.pages ?? -1,
            debugAll: false,
            state: {
                pagination,
                count,
                filters,
            },
            onPaginationChange: setPagination,
            getCoreRowModel: getCoreRowModel(),
            manualPagination: true,
            debugTable: true,
            tableHooks,
        },
    )

    const goToScanReport = (report) => {
        navigate(`/scan/${report.id}`);
    }

    const onFilterChange = (filters) => {
        setFilters(filters);
    }

    return (
        <>
            <ScrollArea sx={{height: "auto"}} onScrollPositionChange={({y}) => setScrolled(y !== 0)}>
                <TableWrapper table={table} scrolled={scrolled} />
            </ScrollArea>

            <Space h="xl"/>
            <Space h="xl"/>

            <PaginationWrapper table={table} count={count} ready={dataQuery.isFetched} />
        </>
    );
};
