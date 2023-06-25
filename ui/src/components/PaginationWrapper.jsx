import React from 'react';
import {createStyles, Flex, Pagination, Select, Text} from "@mantine/core";

const PaginationWrapper = ({ ready, table, count }) => {

    const useStyles = createStyles((theme) => ({}));
    const {theme} = useStyles();

    const getPaginationText = () => {

        if (theme.colorScheme === 'dark') {
            return (
                <>
                    {ready ?
                        <Text style={{color: '#fff' }}>{count} Rows | </Text> :
                        null
                    }
                    <Text style={{color: '#fff' }}>Show</Text>
                </>
            )
        }

        return (
            <>
                {ready ?
                    <Text>{count} Rows | </Text> :
                    null
                }
                <Text>Show</Text>
            </>
        )
    }

    return (
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
            {getPaginationText()}
            <Select w="80px"
                    value={table.getState().pagination.pageSize + ''}
                    onChange={e =>
                        table.setPageSize(Number(e))
                    }
                    data={['10', '20', '50']}/>

        </Flex>
    );
};

export default PaginationWrapper;
