import React from 'react';
import {flexRender} from "@tanstack/react-table";
import {IconArrowDown, IconArrowUp} from "@tabler/icons-react";
import {createStyles, Table} from "@mantine/core";

const TableWrapper = ({scrolled, table}) => {
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

    return (
        <Table striped highlightOnHover withBorder withColumnBorders horizontalSpacing="lg"
               verticalSpacing="xs"
               fontSize="sm"
               sx={{minWidth: 700}}
               // style={{ tableLayout: 'fixed', width: '100%' }}
        >
            <thead className={cx(classes.header, {[classes.scrolled]: scrolled})}>
            {table.getHeaderGroups().map(headerGroup => (
                <tr key={headerGroup.id}>
                    {headerGroup.headers.map(header => {
                        return (
                            <th key={header.id} colSpan={header.colSpan}
                                style={{
                                    width:
                                        header.getSize() !== 150 ? header.getSize() : undefined,
                                }}
                            >
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
    );
};

export default TableWrapper;
