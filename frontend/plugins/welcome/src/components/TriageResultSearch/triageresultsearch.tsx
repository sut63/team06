import React, { FC, useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { Button, TableCell, TextField, } from '@material-ui/core';
import { Content, Header, Page, pageTheme, ContentHeader } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';

import { Cookies } from '../../Cookie';
import Swal from 'sweetalert2';
import { DefaultApi } from '../../api/apis';

import { EntPatient } from '../../api/models/EntPatient';
import { EntTriageResult } from '../../api/models/EntTriageResult';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    table: {
      minWidth: 650,
    },
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    textField: {
      flexWrap: 'wrap',
      width: '50%',
      margin: theme.spacing(1),
    },
    searchbutton: {
      background: '#ff9999',
      color: '#ffffff',
      margin: theme.spacing(1),
    },
    addbutton: {
      background: '#009999',
      color: '#ffffff',
      margin: theme.spacing(1),
    },
    text: {
      color: '#ffffff',
      fontSize: 18,
    }
  }),
);

const TriageResultSearch: FC<{}> = () => {

  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);

  //cookie init
  var cookie = new Cookies();

  //get cookie value
  var nursename = cookie.GetCookie("nursename");

  //variable for query data
  const [triageresult, setTriageresult] = React.useState<EntTriageResult[]>([]);
  const [patientname, setPatientname] = React.useState("");

  //variable for search status
  const [search, setSearch] = useState(false);
  const [status, setStatus] = useState(false);
  const [patients, setPatients] = useState<EntPatient[]>([])

  //Handle value
  const handlePatientname = (event: any) => {
    setSearch(false);
    setStatus(false);
    setPatientname(event.target.value);
  }

  //Aleart function
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: false,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });
  const aleartMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  //Function Check Search Status
  const checksearch = async () => {
    var check = false;
    triageresult.map(item => {
      if (patientname != "") {
        if (item.edges?.patient?.patientName?.startsWith(patientname)) {
          setStatus(true);
          aleartMessage('success', 'พบข้อมูล');
          check = true;
        }
      }
    })
    if (!check) {
      aleartMessage('error', 'ไม่พบข้อมูล');
    }
    if (patientname == "") {
      aleartMessage("info", "แสดงข้อมูลทั้งหมดในระบบ");
    }
  };

  //logout
  const Logout = async () => {
    cookie.SetCookie('nursename', "", 30)
    console.log("Logout success");
    aleartMessage("warning", "ออกจากระบบสำเร็จ!");
    setTimeout(() => {
      window.location.replace("http://localhost:3000/triageresultlogin");
    }, 1200);
  }

  //start list data To Combobox
  useEffect(() => {
    cookie.SetCookie('gotopage', "http://localhost:3000/triageresultsearch", 30);
    //check login status
    if (nursename == "") {
      Swal.fire({
          title: 'เข้าสู่ระบบก่อนใช้งาน',
          position: 'center',
          showConfirmButton: false,
          timer: 3000,
          timerProgressBar: false,
      });
      setTimeout(() => {
          window.location.replace("http://localhost:3000/triageresultlogin");
      }, 3000);
    }

    //list triageresult
    const getTriageResults = async () => {
      const res = await api.listTriageresult({ offset: 0 });
      setTriageresult(res);
    };

    //list patient
    const getPatients = async () => {
      const res = await api.listPatient({ offset: 0 });
      setPatients(res);
    };

    getTriageResults();
    getPatients();
    getTriageResults();
    setLoading(false);

  }, [loading]);

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={'ผลการคัดกรองผู้ป่วย'}
      >
        <TableCell align={
          "center"} >
          <text
            className={classes.text}
          >
            {nursename}
          </text>
          <br /><br />
          <Button
            type="submit"
            variant="contained"
            color="secondary"
            onClick={() => {
              Logout();
            }}
          >
            Logout
        </Button>
        </TableCell>
      </Header>

      <Content>
        <div className={classes.root}>
          <TextField
            className={classes.textField}
            id="symptom"
            name="symptom"
            variant="outlined"
            type="string"
            size="small"
            value={patientname}
            onChange={handlePatientname}
          />

          <Button
            className={classes.searchbutton}
            variant="contained"
            color="inherit"
            onClick={() => {
              checksearch();
              setSearch(true);
            }
            }
          >
            search
          </Button>

          <Button
            className={classes.addbutton}
            component={RouterLink}
            to='/triageresult'
            variant="contained"
          >
            ADD Screening
          </Button>
        </div>

        <br></br>
        <Paper>
          {search ? (
            <div>
              {  status ? (
                <TableContainer component={Paper}>
                  <Table className={classes.table} aria-label="simple table">
                    <TableHead>
                      <TableRow>
                        <TableCell align="center">No.</TableCell>
                        <TableCell align="center">Patient</TableCell>
                        <TableCell align="center">Symptom</TableCell>
                        <TableCell align="center">Height</TableCell>
                        <TableCell align="center">Weight</TableCell>
                        <TableCell align="center">Pressure</TableCell>
                        <TableCell align="center">UrgencyLeval</TableCell>
                        <TableCell align="center">Department</TableCell>
                        <TableCell align="center">Nurse</TableCell>
                        <TableCell align="center">Date</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {triageresult.filter((filter: any) => filter.edges?.patient?.patientName.startsWith(patientname)).map((item: any) => (
                        <TableRow key={item.id}>
                          <TableCell align="center">{item.id}</TableCell>
                          <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                          <TableCell align="center">{item.symptom}</TableCell>
                          <TableCell align="center">{item.height}</TableCell>
                          <TableCell align="center">{item.weight}</TableCell>
                          <TableCell align="center">{item.pressure}</TableCell>
                          <TableCell align="center">{item.edges?.urgencyLevel?.urgencyName}</TableCell>
                          <TableCell align="center">{item.edges?.department?.departmentName}</TableCell>
                          <TableCell align="center">{item.edges?.nurse?.nurseName}</TableCell>
                          <TableCell align="center">{moment(item.triageDate).format('LLLL')}</TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
              )
                : patientname == "" ? (
                  <div>
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                            <TableCell align="center">No.</TableCell>
                            <TableCell align="center">Patient</TableCell>
                            <TableCell align="center">Symptom</TableCell>
                            <TableCell align="center">Height</TableCell>
                            <TableCell align="center">Weight</TableCell>
                            <TableCell align="center">Pressure</TableCell>
                            <TableCell align="center">UrgencyLeval</TableCell>
                            <TableCell align="center">Department</TableCell>
                            <TableCell align="center">Nurse</TableCell>
                            <TableCell align="center">Date</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>
                          {triageresult.sort().map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                              <TableCell align="center">{item.symptom}</TableCell>
                              <TableCell align="center">{item.height}</TableCell>
                              <TableCell align="center">{item.weight}</TableCell>
                              <TableCell align="center">{item.pressure}</TableCell>
                              <TableCell align="center">{item.edges?.urgencyLevel?.urgencyName}</TableCell>
                              <TableCell align="center">{item.edges?.department?.departmentName}</TableCell>
                              <TableCell align="center">{item.edges?.nurse?.nurseName}</TableCell>
                              <TableCell align="center">{moment(item.triageDate).format('LLLL')}</TableCell>
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
      </Content>
    </Page>
  );
};

export default TriageResultSearch;