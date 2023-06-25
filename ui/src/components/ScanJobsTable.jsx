import {
    Box,
    Breadcrumbs, Container,
    createStyles,
    Flex,
    Pagination,
    Paper,
    ScrollArea,
    Select,
    Space,
    Table,
    Text,
    Title
} from "@mantine/core";
import Footer from "./Footer.jsx";
import ScanScoreCell from "./ScanScoreCell.jsx";
import {Link, useNavigate} from "react-router-dom";
import {useQuery} from "@tanstack/react-query";
import {flexRender, getCoreRowModel, useReactTable} from '@tanstack/react-table'
import {Fragment, useMemo, useState} from "react";
import Filters from "../components/Filters.jsx";
import {IconArrowDown, IconArrowUp} from "@tabler/icons-react";
import {API_URL} from "../config/config.js";

export const ScanJobsTable = () => {

    const useStyles = createStyles((theme) => ({
        header: {
            position: 'sticky',
            top: 0,
            backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
            transition: 'box-shadow 150ms ease',

            '&::after': {
                content: '""',
                position: 'absolute',
                left: 0,
                right: 0,
                bottom: 0,
                borderBottom: `1px solid ${
                    theme.colorScheme === 'dark' ? theme.colors.dark[3] : theme.colors.gray[2]
                }`,
            },
        },

        scrolled: {
            boxShadow: theme.shadows.sm,
        },
    }));

    const {classes, cx} = useStyles();
    const [scrolled, setScrolled] = useState(false);

    const navigate = useNavigate();

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
                // body: JSON.stringify(
                //     {
                //         image: image,
                //         score_greater_than: scoreGreaterThan,
                //     }
                // )
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
        },
        {
            header: "Job Id",
            accessorKey: "uuid",
            id: "id",
        },
        {
            header: "Image",
            accessorKey: "image",
            id: "image",
        },
        {
            header: "Vulnerability Count",
            accessorKey: "vulnerabilityCount",
            id: "vulnerability_count",
        },
        // {
        //     header: "Score",
        //     accessorKey: "score",
        //     id: "score",
        //     cell: ScanScoreCell,
        // }

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
                <Table striped highlightOnHover withBorder withColumnBorders horizontalSpacing="lg"
                       verticalSpacing="xs"
                       fontSize="sm"
                       sx={{minWidth: 700}}
                >
                    <thead className={cx(classes.header, {[classes.scrolled]: scrolled})}>
                    {table.getHeaderGroups().map(headerGroup => (
                        <tr key={headerGroup.id}>
                            {headerGroup.headers.map(header => {
                                return (
                                    <th key={header.id} colSpan={header.colSpan}>
                                        {header.isPlaceholder ? null : (
                                            <div>
                                                {flexRender(
                                                    header.column.columnDef.header,
                                                    header.getContext()
                                                )}

                                                {header.isSorted ? (
                                                    header.isSortedDesc ? (
                                                        <IconArrowDown
                                                            name="arrow-down"
                                                            size="sm"
                                                            color="primary"
                                                            ml="xs"
                                                        />
                                                    ) : (
                                                        <IconArrowUp
                                                            name="arrow-up"
                                                            size="sm"
                                                            color="primary"
                                                            ml="xs"
                                                        />
                                                    )
                                                ) : null}
                                            </div>
                                        )}
                                    </th>
                                )
                            })}
                        </tr>
                    ))}
                    </thead>
                    <tbody>
                    {table.getRowModel().rows.map(row => {
                        return (
                            <tr key={row.id} style={{cursor: "pointer"}}>
                                {row.getVisibleCells().map(cell => {
                                    return (
                                        <td key={cell.id}>
                                            {flexRender(
                                                cell.column.columnDef.cell,
                                                cell.getContext()
                                            )}
                                        </td>
                                    )
                                })}
                            </tr>
                        )
                    })}
                    </tbody>
                </Table>
            </ScrollArea>

            <Space h="xl"/>
            <Space h="xl"/>

            <Flex
                mb="20px"
                mih={50}
                gap="md"
                justify="flex-end"
                align="center"
                direction="row"
                wrap="wrap"
            >
                <Pagination position="center" page={table.getState().pagination.pageIndex + 1}
                            onChange={(i) => table.setPageIndex(i - 1)} total={table.getPageCount()}/>
                {dataQuery.isFetched ?
                    <Text>{count} Rows | </Text> :
                    null
                }
                <Text>Show</Text>
                <Select w="80px"
                        value={table.getState().pagination.pageSize + ''}
                        onChange={e =>
                            table.setPageSize(Number(e))
                        }
                        data={['10', '20', '50']}/>

            </Flex>
        </>
    );
};
