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
//icon
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import DeleteTwoToneIcon from '@material-ui/icons/DeleteTwoTone';

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

  const [checkpersonalname, setPersonalnames] = useState(false);
  const [personal, setPersonal] = useState<EntPersonal[]>([])

  const [personalname, setPersonalname] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getPersonals = async () => {
      const res = await http.listPersonal({ offset: 0 });
      setLoading(false);
      setPersonal(res);
    };
    getPersonals();
  }, [loading]);

  const personalnamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setPersonalnames(false);
    setPersonalname(event.target.value as string);

  };

  const cleardata = () => {
    setPersonalname("");
    setSearch(false);
    setPersonalnames(false);
    setSearch(false);

  }

  const checkresearch = async () => {
    var check = false;
    personal.map(item => {
      if (personalname != "") {
        if (item.personalname?.includes(personalname)) {
          setPersonalnames(true);
          alertMessage("success", "ค้นหาข้อมูลบุคลากรสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ค้นหาข้อมูลบุคลากรไม่สำเร็จ");
    }
    console.log(checkpersonalname)
    if (personalname == "") {
      alertMessage("info", "แสดงข้อมูลบุคลากรทั้งหมดทั้งหมดในระบบ");
    }
  };

  return (

    <Page theme={pageTheme.tool}>
      <Header  title="ระบบข้อมูลบุคลากร" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
        <Button 
          style={{ marginLeft: 20 }} 
          href="/"
          variant="contained"  
          startIcon={<ExitToAppIcon/>}
          > 
          ออกจากระบบ 
        </Button>
      </Header>
      <Content>
        <ContentHeader title="เพิ่มข้อมูลบุคลากร">
        <div>&nbsp;&nbsp;&nbsp;</div>
          <Button  
          onClick={() => {
            checkresearch();
            setSearch(true);
          }}
          variant="contained" 
          color="secondary" 
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
          > 
          เคลียร์ข้อมูล 
          </Button>
        <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/Personalwelcome" 
          variant="contained"
          color="primary"
          startIcon={<CancelTwoToneIcon/>}
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
            <div className={classes.paper}><strong>กรอก "ชื่อ" เพื่อทำการค้นหา</strong></div>
            <TextField
            style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:5}}
            className={classes.textField}
              id="personalname"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={personalname}
              onChange={personalnamehandlehange}
            />
            </FormControl>
            </form>
            </div>
        
        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              {search ? (
                <div>
                  {  checkpersonalname ? (
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

                          {personal.filter((filter: any) => filter.personalname.includes(personalname)).map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.edges?.title?.titlename}</TableCell>
                              <TableCell align="center">{item.personalname}</TableCell>
                              <TableCell align="center">{item.edges?.gender?.gendername}</TableCell>
                              <TableCell align="center">{item.email}</TableCell>
                              <TableCell align="center">{item.edges?.department?.departmentname}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : personalname == "" ? (
                      <div>
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
                                  <TableCell align="center">{item.edges?.title?.titlename}</TableCell>
                                  <TableCell align="center">{item.personalname}</TableCell>
                                  <TableCell align="center">{item.edges?.gender?.gendername}</TableCell>
                                  <TableCell align="center">{item.email}</TableCell>
                                  <TableCell align="center">{item.edges?.department?.departmentname}</TableCell>
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