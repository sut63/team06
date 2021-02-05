import React, { useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, TextField, Button, Grid } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import Swal from 'sweetalert2'
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';

import { EntPatient } from '../../api/models/EntPatient';
import { EntAppointmentResults } from '../../api/models/EntAppointmentResults';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(2),
    },
    insideLabel: {
      margin: theme.spacing(1),
    },
    button: {
      marginLeft: '50px',
    },
    textField: {
      width: 600,
      marginLeft: 10,
      marginRight: -10,
    },
    select: {
      width: 600,
      marginLeft: 10,
      marginRight: -10,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    pa: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
  }),
);

const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});


export default function ComponentsTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  const [checkvalue, setCheckValue] = useState(false);

  const [checkHospitalNumberID, setHospitalNumberIDs] = useState(false);
  const [patients, setPatients] = useState<EntPatient[]>([])
  const [appointmentresultss, setAppointmentResultss] = useState<EntAppointmentResults[]>([])

  const [hospitalnumberID, sethospitalnumberID] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getPatients = async () => {
      const res = await http.listPatient({ offset: 0 });
      setLoading(false);
      setPatients(res);
    };
    getPatients();
    const getAppointmentResultss = async () => {
        const res = await http.listAppointmentresults({ offset: 0 });
        setLoading(false);
        setAppointmentResultss(res);
      };
      getAppointmentResultss();
  }, [loading]);

  const hospitalnumberIDhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setHospitalNumberIDs(false);
    sethospitalnumberID(event.target.value as string);

  };

  const cleardata = () => {
    sethospitalnumberID("");
    setSearch(false);
    setHospitalNumberIDs(false);
    setSearch(false);

  }

  const checkresearch = async () => {
    var check = false;
    appointmentresultss.map(item => {
      if (hospitalnumberID != "") {
        if (item.edges?.appointmentResultsToPatient?.hospitalNumber == hospitalnumberID) {
            setHospitalNumberIDs(true);
          alertMessage("success", "ค้นหาข้อมูลการนัดหมายของผู้ป่วยสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลการนัดหมายของผู้ป่วยที่ค้นหา");
    }
    console.log(checkHospitalNumberID)
    if (hospitalnumberID == "") {
      alertMessage("info", "แสดงข้อมูลการนัดหมายทั้งหมดในระบบ");
    }
  };

  return (

    <Page theme={pageTheme.home}>
      <Header title="ระบบค้นหาข้อมูลการนัดหมาย">
      </Header>
      <Content>
        <ContentHeader title="ค้นหาข้อมูลการนัดหมายของผู้ป่วย">
          <Button
            href="/createappointment"
            variant="contained"
            color="primary"
          >
            กลับสู่ระบบนัดหมาย
          </Button>
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
              size="small"
            >
              <div className={classes.paper}><strong>กรอก "เลขประจำตัวผู้ป่วย" เพื่อทำการค้นหา</strong></div>
              <TextField
                style={{ width: 500, marginLeft: 7, marginRight: -7, marginTop: 5 }}
                className={classes.textField}
                id="searchAppoint"
                variant="outlined"
                color="primary"
                type="string"
                size="small"
                value={hospitalnumberID}
                onChange={hospitalnumberIDhandlehange}
              />
            </FormControl>
          </form>
        </div>
        <div className={classes.root}>
          <Button
            onClick={() => {
              checkresearch();
              setSearch(true);
            }}
            variant="contained"
            color="secondary"
            startIcon={<SearchTwoToneIcon />}
          >
            ค้นหาข้อมูล
          </Button>
          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button
            onClick={() => {
              cleardata();
            }}
            variant="contained"
          >
            เคลียร์ข้อมูล
          </Button>
        </div>

        <br></br>

        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              {search ? (
                <div>
                  {  checkHospitalNumberID ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                            <TableCell align="center">ลำดับ</TableCell>
                            <TableCell align="center">ผู้บันทึก</TableCell>
                            <TableCell align="center">ผู้ป่วย</TableCell>
                            <TableCell align="center">แพทย์</TableCell>
                            <TableCell align="center">ห้องตรวจ</TableCell>
                            <TableCell align="center">สาเหตุ</TableCell>
                            <TableCell align="center">คำแนะนำ</TableCell>
                            <TableCell align="center">ยื่นบัตรนัด<br></br>ก่อนเวลานัด<br></br><br></br>ชั่วโมง</TableCell>
                            <TableCell align="left"><br></br><br></br><br></br>นาที</TableCell>
                            <TableCell align="center">วันนัดหมาย</TableCell>
                            <TableCell align="center">เวลานัดหมาย</TableCell>
                            <TableCell align="center">วัน/เวลาที่บันทึก</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {appointmentresultss.filter((filter: any) => filter.edges?.appointmentResultsToPatient?.hospitalNumber.includes(hospitalnumberID)).map((item: any) => (
                            <TableRow key={item.id}>
                                <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">{item.edges?.appointmentResultsToNurse?.nurseName}</TableCell>
                                <TableCell align="center">{item.edges?.appointmentResultsToPatient?.hospitalNumber}</TableCell>
                                <TableCell align="center">{item.edges?.appointmentResultsToDoctor?.doctorName}</TableCell>
                                <TableCell align="center">{item.edges?.appointmentResultsToRoom?.roomName}</TableCell>

                                <TableCell align="center">{item.causeAppoint}</TableCell>
                                <TableCell align="center">{item.advice}</TableCell>
                                <TableCell align="center">{item.hourBeforeAppoint||0} hr</TableCell>
                                <TableCell align="left">{item.minuteBeforeAppoint||0} m</TableCell>

                                <TableCell align="center">{moment(item.dateAppoint).format("ll")}</TableCell>
                                <TableCell align="center">{moment(item.timeAppoint).format("LT")}</TableCell>
                                <TableCell align="center">{moment(item.addtimeSave).format("lll")}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : hospitalnumberID == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                                <TableCell align="center">ลำดับ</TableCell>
                                <TableCell align="center">ผู้บันทึก</TableCell>
                                <TableCell align="center">ผู้ป่วย</TableCell>
                                <TableCell align="center">แพทย์</TableCell>
                                <TableCell align="center">ห้องตรวจ</TableCell>
                                <TableCell align="center">สาเหตุ</TableCell>
                                <TableCell align="center">คำแนะนำ</TableCell>
                                <TableCell align="center">ยื่นบัตรนัด<br></br>ก่อนเวลานัด<br></br><br></br>ชั่วโมง</TableCell>
                                <TableCell align="left"><br></br><br></br><br></br>นาที</TableCell>
                                <TableCell align="center">วันนัดหมาย</TableCell>
                                <TableCell align="center">เวลานัดหมาย</TableCell>
                                <TableCell align="center">วัน/เวลาที่บันทึก</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {appointmentresultss.map((item: any) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.appointmentResultsToNurse?.nurseName}</TableCell>
                                    <TableCell align="center">{item.edges?.appointmentResultsToPatient?.hospitalNumber}</TableCell>
                                    <TableCell align="center">{item.edges?.appointmentResultsToDoctor?.doctorName}</TableCell>
                                    <TableCell align="center">{item.edges?.appointmentResultsToRoom?.roomName}</TableCell>

                                    <TableCell align="center">{item.causeAppoint}</TableCell>
                                    <TableCell align="center">{item.advice}</TableCell>
                                    <TableCell align="center">{item.hourBeforeAppoint||0} hr</TableCell>
                                    <TableCell align="left">{item.minuteBeforeAppoint||0} m</TableCell>

                                    <TableCell align="center">{moment(item.dateAppoint).format("ll")}</TableCell>
                                    <TableCell align="center">{moment(item.timeAppoint).format("LT")}</TableCell>
                                    <TableCell align="center">{moment(item.addtimeSave).format("lll")}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : 
                  <TableContainer component={Paper}>
                  <Table className={classes.table} aria-label="simple table">
                    <TableHead>
                      <TableRow>
                        <TableCell align="center">ลำดับ</TableCell>
                        <TableCell align="center">ผู้บันทึก</TableCell>
                        <TableCell align="center">ผู้ป่วย</TableCell>
                        <TableCell align="center">แพทย์</TableCell>
                        <TableCell align="center">ห้องตรวจ</TableCell>
                        <TableCell align="center">สาเหตุ</TableCell>
                        <TableCell align="center">คำแนะนำ</TableCell>
                        <TableCell align="center">ยื่นบัตรนัด<br></br>ก่อนเวลานัด<br></br><br></br>ชั่วโมง</TableCell>
                        <TableCell align="left"><br></br><br></br><br></br>นาที</TableCell>
                        <TableCell align="center">วันนัดหมาย</TableCell>
                        <TableCell align="center">เวลานัดหมาย</TableCell>
                        <TableCell align="center">วัน/เวลาที่บันทึก</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>

                      {appointmentresultss.map((item: any) => (
                        <TableRow key={item.id}>
                            <TableCell align="center">{item.id}</TableCell>
                            <TableCell align="center">{item.edges?.appointmentResultsToNurse?.nurseName}</TableCell>
                            <TableCell align="center">{item.edges?.appointmentResultsToPatient?.hospitalNumber}</TableCell>
                            <TableCell align="center">{item.edges?.appointmentResultsToDoctor?.doctorName}</TableCell>
                            <TableCell align="center">{item.edges?.appointmentResultsToRoom?.roomName}</TableCell>

                            <TableCell align="center">{item.causeAppoint}</TableCell>
                            <TableCell align="center">{item.advice}</TableCell>
                            <TableCell align="center">{item.hourBeforeAppoint||0} hr</TableCell>
                            <TableCell align="left">{item.minuteBeforeAppoint||0} m</TableCell>

                            <TableCell align="center">{moment(item.dateAppoint).format("ll")}</TableCell>
                            <TableCell align="center">{moment(item.timeAppoint).format("LT")}</TableCell>
                            <TableCell align="center">{moment(item.addtimeSave).format("lll")}</TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
                }
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}