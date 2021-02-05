import React, { useState, useEffect, FC } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import moment from 'moment';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';

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
import { EntPatient } from '../../api/models/EntPatient';

const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
});

const PatientRegistrationTables: FC<{}> = () => {

    const classes = useStyles();
    const http = new DefaultApi();
    const [patients, setPatients] = useState<EntPatient[]>([]);

    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getPatients = async () => {
            const res = await http.listPatient({ limit: 10, offset: 0 });
            setLoading(false);
            setPatients(res);
        };
        getPatients();
    }, [loading]);

    const deletePatients = async (id: number) => {
        const res = await http.deletePatient({ id: id });
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
                    <Button
                        href="/patientregistration"
                        variant="contained"
                        color="primary"
                    >
                        ลงทะเบียนผู้ป่วย
                    </Button>
                         &nbsp;&nbsp;&nbsp;
                    <Button
                        href="/searchpatient"
                        variant="contained"
                        color="primary"
                        startIcon={<SearchTwoToneIcon />}
                    >
                        ค้นหาข้อมูลผู้ป่วย
                    </Button>

                </ContentHeader>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">เลขประจำตัวประชาชน</TableCell>
                                <TableCell align="center">คำนำหน้า</TableCell>
                                <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                                <TableCell align="center">เพศ</TableCell>
                                <TableCell align="center">หมู่เลือด</TableCell>
                                <TableCell align="center">หมายเลขผู้ป่วย</TableCell>
                                <TableCell align="center">ประวัติการแพ้ยา</TableCell>
                                <TableCell align="center">เบอร์โทรศัพท์</TableCell>
                                <TableCell align="center">วันที่ลงทะเบียน</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {patients.map((item: any) => (

                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.personalID}</TableCell>
                                    <TableCell align="center">{item.edges?.prefix?.prefixValue}</TableCell>
                                    <TableCell align="center">{item.patientName}</TableCell>
                                    <TableCell align="center">{item.edges?.gender?.genderValue}</TableCell>
                                    <TableCell align="center">{item.edges?.bloodtype?.bloodValue}</TableCell>
                                    <TableCell align="center">{item.hospitalNumber}</TableCell>
                                    <TableCell align="center">{item.drugAllergy}</TableCell>
                                    <TableCell align="center">{item.mobileNumber}</TableCell>
                                    <TableCell align="center">{moment(item.added).format('LLL')}</TableCell>
                                    <TableCell align="center">
                                        <Button
                                            onClick={() => {
                                                deletePatients(item.id);
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
export default PatientRegistrationTables;
