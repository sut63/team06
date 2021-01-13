import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import { TableCell, } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import {
    Content,
    Header,
    Page,
    pageTheme,

    Link,
} from '@backstage/core';
const Medicalresults: FC<{}> = () => {
    const profile = { givenName: 'Medical Procedure' };
    return (
        <Page theme={pageTheme.home}>
            <Header
                title={` ${profile.givenName || 'Laboratory Data'}`}
                subtitle="ยินดีต้อนรับเข้าสู่ระบบบันทึกการทำหัตถการ"
            >
                <Link component={RouterLink} to="/medicalprocedure">
                    <br /><br />
                    <Button size="large" variant="contained" color="secondary">
                        ADD
                    </Button>
                </Link>
            </Header>
            
            <Content>


                <TableCell></TableCell>
                
                <ComponanceTable></ComponanceTable>

            </Content>
        </Page>
    );
};
export default Medicalresults;