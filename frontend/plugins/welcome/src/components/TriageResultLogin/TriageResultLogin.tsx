import React, { FC, useEffect, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { DefaultApi } from '../../api/apis';
import Swal from 'sweetalert2';
import { Cookies } from '../../Cookie';

import { EntNurse } from '../../api/models/EntNurse';

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%',
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignIn: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [path, setPath] = React.useState("");

  //cookie
  var cookie = new Cookies();

  //query
  const [nurses, setNurses] = React.useState<EntNurse[]>([]);

  //input
  const [username, setUsername] = React.useState(String);
  const [password, setPassword] = React.useState(String);

  //HandleValue
  const usernameHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUsername(event.target.value as string);
  };
  const passwordHandle = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPassword(event.target.value as string);
  };

  //Aleart
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

  const aleartMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  //start
  useEffect(() => {
    const getNurses = async () => {
      const res = await api.listNurse({ limit: 1000, offset: 0 });
      setNurses(res);
    };
    getNurses();
  }, [loading]);

  const Login = async () => {
    var status = false;
    console.log(username);
    console.log(password);
    nurses.map((item: EntNurse) => {
      if (item.nurseUsername === username && item.nursePassword === password) {
        cookie.SetCookie('nursename', item.nurseName, 30);
        cookie.SetCookie('nurseID', item.id, 30);
        aleartMessage("success", "เข้าสู่ระบบสำเร็จ");
        status = true;
      }
    });
    if (status) {
      window.location.replace("http://localhost:3000/triageresult");
    }
    else {
      aleartMessage("error", "เข้าสู่ระบบไม่สำเร็จ");
      setPath("/triageresultlogin");
    }
  }

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="username"
            label="Username"
            name="username"
            autoComplete="email"

            value={username}
            onChange={usernameHandle}
            autoFocus
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            value={password}
            onChange={passwordHandle}
          />

          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={() => {
              Login();
            }}
            component={RouterLink}
            to={path}
          >
            Sign In
          </Button>

          <Grid container>
          </Grid>
        </form>
      </div>
      <Box mt={8}>
      </Box>
    </Container>
  );
};

export default SignIn;
