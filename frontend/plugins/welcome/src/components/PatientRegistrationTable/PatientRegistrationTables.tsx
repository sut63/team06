import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';

import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
} from '@backstage/core';
  
import {
    Box,
  } from '@material-ui/core';

import { DefaultApi } from '../../api/apis';
import { EntPatientDetail } from '../../api/models/EntPatientDetail';

const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
});

export default function ComponentsTable() {
    const classes = useStyles();
    const http = new DefaultApi();
    const [patientDetails, setPatientDetails] = useState<EntPatientDetail[]>([]);

    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getPatients = async () => {
            const res = await http.listPatientDetail({ limit: 10, offset: 0 });
            setLoading(false);
            setPatientDetails(res);
        };
        getPatients();
    }, [loading]);

    const deletePatientDetails = async (id: number) => {
        const res = await http.deletePatientdetail({ id: id });
        setLoading(true);
    };

    return (
        <Page theme={pageTheme.home}>
            <Header
                title="ระบบลงทะเบียนผู้ป่วย"
                subtitle="สามารถดูผลลงทะเบียนผู้ป่วยได้ที่นี่"
            ></Header>

            <Content>
                <ContentHeader title="ผลการลงทะเบียนผู้ป่วย">
                    <Box
                        display="flex"
                        justifyContent="flex-start"
                        flexDirection="column"
                    >
                    </Box>

                </ContentHeader>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">คำนำหน้า</TableCell>
                                <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                                <TableCell align="center">เพศ</TableCell>
                                <TableCell align="center">หมู่เลือด</TableCell>
                                <TableCell align="center">หมายเลขผู้ป่วย</TableCell>
                                <TableCell align="center">ประวัติการแพ้ยา</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {patientDetails.map((item: any) => (

                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges?.prefix?.prefixValue}</TableCell>
                                    <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                                    <TableCell align="center">{item.edges?.gender?.genderValue}</TableCell>
                                    <TableCell align="center">{item.edges?.bloodtype?.bloodValue}</TableCell>
                                    <TableCell align="center">{item.edges?.patient?.hospitalNumber}</TableCell>
                                    <TableCell align="center">{item.edges?.patient?.drugAllergy}</TableCell>
                                    <TableCell align="center">
                                        <Button
                                            onClick={() => {
                                                deletePatientDetails(item.id);
                                            }}
                                            style={{ marginLeft: 10 }}
                                            variant="contained"
                                            color="secondary"
                                        >
                                            Delete
                                    </Button>
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Content>
        </Page>
    );
}