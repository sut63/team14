//style
import React, {  useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, InputAdornment, Grid } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
//api
import { DefaultApi } from '../../api/apis';
//table
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
//entity
import { EntPersonal } from '../../api/models/EntPersonal';
//alert
import Swal from 'sweetalert2'
import { Alert } from '@material-ui/lab';
//icon
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import DeleteTwoToneIcon from '@material-ui/icons/DeleteTwoTone';

import { styled } from '@material-ui/core/styles';
import { compose, spacing, palette, sizing, shadows   } from '@material-ui/system';

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
   marginLeft: '40px',
  },
  textField: {
   width: 500 ,
   marginLeft:7,
   marginRight:-7,
  },
  select: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
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
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);

  const [checkpersonalname, setPersonalnames] = useState(false);
  const [personal, setPersonal] = useState<EntPersonal[]>([])

  const [pname, setpname] = useState(String);

  const Box = styled('div')(compose(spacing, palette, shadows, sizing ));

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  const personalnamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setpname(event.target.value as string);
  };

  const cleardata = () => {
    setpname("");
    setStatus(false)
    setPersonal([]);
  }

  const SearchPersonal = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchpersonals?personal=${pname}`;
    const requestOptions = {
      method: 'GET',
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setStatus(true);
        setAlert(false);
        setPersonal([]);
        if (data.data != null) {
          if (data.data.length >= 1) {
            setStatus(true);
            setAlert(true);
            console.log(data.data)
            setPersonal(data.data);
          }
        }
      });
    }

  return (

    <Page theme={pageTheme.tool}>
      <Header  title="ระบบข้อมูลบุคลากร" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button  
          onClick={() => {
            SearchPersonal();
          }}
          variant="contained" 
          style={{background:"#10772C",color:"white"}}
          startIcon={<SearchTwoToneIcon/>}
          > 
          ค้นหาข้อมูล 
          </Button>
          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button  
          onClick={() => {
            cleardata();
          }}
          variant="contained"  
          startIcon={<DeleteTwoToneIcon/>}
          style={{background:"#F4770B",color:"white"}}
          > 
          เคลียร์ข้อมูล 
          </Button>
          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/Personaltable" 
          variant="contained"
          style={{background:"#F62915",color:"white"}}
          startIcon={<CancelTwoToneIcon/>}
          > 
          ย้อนกลับ 
          </Button>
          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          style={{background:"#B9BAC5"}}
          href="/"
          variant="contained"  
          startIcon={<ExitToAppIcon/>}
          > 
          ออกจากระบบ 
          </Button>
      </Header>
      <Content>
        <ContentHeader title="ค้นหาข้อมูลบุคลากร">
          {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    ✔️ ค้นหาข้อมูลบุคลากรสำเร็จ ✔️
                  </Alert>
                )
                  : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ❌ ค้นหาข้อมูลบุคลากรไม่สำเร็จ ❌
                    </Alert>
                  )}
              </div>
            ) : null}
        </ContentHeader>

        <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
            size="small"
          >
            <div className={classes.paper}><strong>กรอก "ชื่อ" เพื่อทำการค้นหา</strong></div>
            <TextField
            style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:5}}
            className={classes.textField}
              id="personalname"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={pname}
              onChange={personalnamehandlehange}
            />
            </FormControl>
            </form>
            </div>
        
        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="simple table">
                   <TableHead>
                      <TableRow>
                        <TableCell align="center">No</TableCell>
                        <TableCell align="center">คำนำหน้าชื่อ</TableCell>
                        <TableCell align="center">ชื่อ-นามสกุล</TableCell>
                        <TableCell align="center">เพศ</TableCell>
                        <TableCell align="center">อีเมลล์</TableCell>
                        <TableCell align="center">แผนก</TableCell>
                        </TableRow>
                      </TableHead>
                    <TableBody>
                    {personal.map((item: any) => (
                      <TableRow key={item.id}>
                        <TableCell align="center">{item.id}</TableCell>
                        <TableCell align="center">{item.edges?.Title?.titlename}</TableCell>
                        <TableCell align="center">{item.Personalname}</TableCell>
                        <TableCell align="center">{item.edges?.Gender?.Gendername}</TableCell>
                        <TableCell align="center">{item.Email}</TableCell>
                        <TableCell align="center">{item.edges?.Department?.Departmentname}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}