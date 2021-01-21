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
import { EntDiagnosis } from '../../api/models/EntDiagnosis';

const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
});

export default function ComponentsTable() {
    const classes = useStyles();
    const http = new DefaultApi();
    const [diagnosiss, setDiagnosiss] = useState<EntDiagnosis[]>([]);

    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getDiagnosiss = async () => {
            const res = await http.listDiagnosis({ limit: 10, offset: 0 });
            setLoading(false);
            setDiagnosiss(res);
        };
        getDiagnosiss();
    }, [loading]);

    const deleteDiagnosiss = async (id: number) => {
        const res = await http.deleteDiagnosis({ id: id });
        setLoading(true);
    };

    return (
        <Page theme={pageTheme.home}>
            <Header
                title="ระบบตรวจวินิจฉัย"
                subtitle="สามารถดูผลลงตรวจวินิจฉัยที่นี่"
            ></Header>

            <Content>
                <ContentHeader title="ผลการตรวจวินิจฉัย">
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
                                <TableCell align="center">ชื่อแพทย์</TableCell>
                                <TableCell align="center">ชื่อผู้ป่วย </TableCell>
                                <TableCell align="center">ประเภทการรักษา</TableCell>
                                <TableCell align="center">อาการ</TableCell>
                                <TableCell align="center">ความคิดเห็นจากแพทย์</TableCell>
                                <TableCell align="center">หมายเหตุ</TableCell>
                                <TableCell align="center">วันที่</TableCell>
                                
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {diagnosiss.map((item: any) => (

                                <TableRow key={item.id}>
                                    
                                    <TableCell align="center">{item.edges?.Doctor?.doctorName}</TableCell>
                                    <TableCell align="center">{item.edges?.patients?.patientName}</TableCell>
                                    <TableCell align="center">{item.edges?.treatmentype?.Type}</TableCell>
                                    <TableCell align="center">{item.symptom}</TableCell>
                                    <TableCell align="center">{item.opinionresult}</TableCell>
                                    <TableCell align="center">{item.note}</TableCell>
                                    <TableCell align="center">{item.diagnosisDate}</TableCell>
                                    <TableCell align="center">
                                        <Button
                                            onClick={() => {
                                                deleteDiagnosiss(item.id);
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
