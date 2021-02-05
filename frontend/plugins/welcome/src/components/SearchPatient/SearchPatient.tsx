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

  const [checkpersonalID, setPersonalIDs] = useState(false);
  const [patients, setPatients] = useState<EntPatient[]>([])

  const [personalID, setPersonalID] = useState(String);
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
  }, [loading]);

  const personalIDhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setPersonalIDs(false);
    setPersonalID(event.target.value as string);

  };

  const cleardata = () => {
    setPersonalID("");
    setSearch(false);
    setPersonalIDs(false);
    setSearch(false);

  }

  const checkresearch = async () => {
    var check = false;
    patients.map(item => {
      if (personalID != "") {
        if (item.personalID?.includes(personalID)) {
          setPersonalIDs(true);
          alertMessage("success", "ค้นหาข้อมูลผู้ป่วยสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลผู้ป่วยที่ค้นหา");
    }
    console.log(checkpersonalID)
    if (personalID == "") {
      alertMessage("info", "แสดงข้อมูลผู้ป่วยทั้งหมดทั้งหมดในระบบ");
    }
  };

  return (

    <Page theme={pageTheme.home}>
      <Header title="ระบบค้นหาข้อมูลผู้ป่วย">
      </Header>
      <Content>
        <ContentHeader title="ดูข้อมูลผู้ป่วย">
          <Button
            href="/patientregistration"
            variant="contained"
            color="primary"
          >
            ลงทะเบียนผู้ป่วย
          </Button>
          &nbsp;&nbsp;&nbsp;
          <Button
            href="/patientregistrationtable"
            variant="contained"
            color="primary"
          >
            ย้อนกลับ
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
              <div className={classes.paper}><strong>กรอก "เลขประจำตัวประชาชน" เพื่อทำการค้นหา</strong></div>
              <TextField
                style={{ width: 500, marginLeft: 7, marginRight: -7, marginTop: 5 }}
                className={classes.textField}
                id="personalID"
                variant="outlined"
                color="primary"
                type="string"
                size="small"
                value={personalID}
                onChange={personalIDhandlehange}
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
                  {  checkpersonalID ? (
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

                          {patients.filter((filter: any) => filter.personalID.includes(personalID)).map((item: any) => (
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
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : personalID == "" ? (
                      <div>
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
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}