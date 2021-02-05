import React, { FC, useState, useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { DefaultApi, EntNurse } from '../../api';
import { Alert } from '@material-ui/lab'; // alert
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Swal from 'sweetalert2';
const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: '#C7B2DE	',
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

  const [status, SetStatus] = useState(false);
  const [loading, SetLoading] = useState(true);
  const [alert, SetAlert] = useState(Boolean);

  const [nurse, setNurse] = useState<EntNurse[]>([]);
  const [username, setUsername] = useState(String);
  const [password, setPassword] = useState(String);

  const PasswordhandelChange = (event: any) => {
    setPassword(event.target.value as string);
  };
  const UsernamehandelChange = (event: any) => {
    setUsername(event.target.value as string);
  };
  console.log('username', username);

  useEffect(() => {
    const getNurse = async () => {
      const res: any = await api.listNurse({ offset: 0 });
      SetLoading(false);
      setNurse(res);
    };
    getNurse();
    localStorage.clear();
  }, [loading]);

  const SigninhandleChange = async () => {
    nurse.map((item: EntNurse) => {
      console.log(item.nurseUsername);
      if (item.nurseUsername == username && item.nursePassword == password) {
        SetAlert(true);
        localStorage.setItem('nurse-id', JSON.stringify(item.id));
        localStorage.setItem('nurse-nurseName', JSON.stringify(item.nurseName));
        localStorage.setItem('nurse-nurseUsername', JSON.stringify(item.nurseUsername));
        history.pushState('', '', '/Create');
      }
    });

    SetStatus(true);
    const timer = setTimeout(() => {
      SetStatus(false);
    }, 99999999);
  };

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      {status ? (
        <div>
          {alert ? (
              <Alert severity="success">Login Succese</Alert>
               ) 
            : (
              <Alert severity="error" style={{ marginTop: 50 }}>
                email or password incorrect!!!
              </Alert>
          )}
        </div>
      ) : null}
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
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            onChange={UsernamehandelChange}
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
            onChange={PasswordhandelChange}
          />
          <Button
            style={{ backgroundColor: '#C7B2DE	' }}
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={() => {
              SigninhandleChange();
            }}

          >
            Sign In
          </Button>
        </form>
      </div>
      <Box mt={8}>
      </Box>
    </Container>
  );
};

export default SignIn;