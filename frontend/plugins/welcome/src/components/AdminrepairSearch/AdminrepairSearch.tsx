import React, { useEffect,useState } from 'react';
import { ContentHeader,Content, Header, Page, pageTheme } from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
  InputLabel,
  TextField,
}from '@material-ui/core';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import { DefaultApi } from '../../api/apis';
import { EntAdminrepair } from '../../api/models/EntAdminrepair';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import Button from '@material-ui/core/Button';
import Swal from 'sweetalert2';
import SearchIcon from '@material-ui/icons/Search';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import FormHelperText from '@material-ui/core/FormHelperText';
import Input from '@material-ui/core/Input';
import RestoreFromTrashIcon from '@material-ui/icons/RestoreFromTrash';

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 20,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
    marginRight: theme.spacing(-1),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
  table: {
    minWidth: 500,
  },
}));


  export default function AdminrepairCreate(){
    const classes = useStyles();
    const api = new DefaultApi();
    const [loading, setLoading] = useState(true);
    const [searchall, setSearchall] = useState(true);
    const [checktext, setChecktext] = useState(true);
    const [status , setStatus] = useState(true);
  
    const [searchadminrepair , setSearchadminrepair] = useState(false);
    const [adminrepair, setAdminrepair] = useState<EntAdminrepair[]>([])
    const [adminrepairTable, setAdminrepairTable] = useState<EntAdminrepair[]>([])
  
    const [numberrepair, setNumberrepair] = useState(String);
  
    const [numberrepairError, setNumberrepairError] = useState('');


  
    useEffect(() =>{
        const getAdminrepairs = async () => {
            const res = await api.listAdminrepair({ offset: 0 });
            setLoading(false);
            setAdminrepairTable(res);
          };
          getAdminrepairs();
    },[loading]);
  
    const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 5000,
      timerProgressBar: true,
      didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
      },
    });
  
    const alertMessage = (icon: any, title: any) => {
      Toast.fire({
        icon: icon,
        title: title,
      });
    }
  
    const validareNumberrepair = (val: string)=>{
      return val.match("[A]+[M]+[P]+[-]+[0-9]+[0-9]") && val.length == 6;
    }
    const checkPattern = (id:string,value:string)=>{
      switch(id) {
        case 'numberrepair':
          validareNumberrepair(value) ? setNumberrepairError("") : setNumberrepairError("รูปแบบหมายเลข AMP-XX");
          return;
        default:
          return;
      }
    }
  
    const Numberrepairhandlechange = (event: React.ChangeEvent<{value: any }>) => {
      const { value } = event.target;
      const validateValue = value
      checkPattern('numberrepair', validateValue)
      setNumberrepair(event.target.value as string);
    }
  
    const SearchAdminrepair = async () => {
      const apiUrl = `http://localhost:8080/api/v1/searchadminrepairs?adminrepair=${numberrepair}`;
      const requestOptions = {
        method: 'GET',
      };
      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data.data)
          setAdminrepair([]);
          if (data.data != null) {
            if (data.data.length >= 1) {
              setSearchadminrepair(true);
              setSearchall(false);
              setChecktext(false);
              setStatus(false);
              console.log(data.data)
              setAdminrepair(data.data);
              alertMessage("success", "ค้นหาหมายเลขบันทึกซ่อมแซมคอมพิวเตอร์ของพนักงาน สำเร็จ!!");
              if(!status){
                alertMessage("error", "อยู่ระหว่างการค้นหา ไม่สามารถค้นหาหมายเลขอื่นได้ กรุณาลบข้อมูลเก่าออกก่อน แล้วโปรลองใหม่!!");
              }
            }else{
                alertMessage("error", "ค้นหาหมายเลขบันทึกซ่อมแซมคอมพิวเตอร์ของพนักงาน ไม่พบ!!");
            }
          }else{
            setSearchall(true);
            setSearchadminrepair(false);
            alertMessage("error", "กรุณาใส่จำนวนหมายเลขให้ถูกก่อนกดค้นหา!!");
          }
        });
      }

    const cleardata = () => {
        setNumberrepair("");
        setSearchadminrepair(false);
        setStatus(true);
        setSearchall(true);
        setChecktext(true);
    }
  
    return (
      <Page theme={pageTheme.home}>
      <Header title="ระบบค้นหาบันทึกซ่อมแซมคอมพิวเตอร์ของพนักงาน" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
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
           <ContentHeader title="ค้นหาหมายเลข บันทึกซ่อมแซมคอมพิวเตอร์ของพนักงาน">{checktext ? (<TextField 
              style={{  width: 150,marginRight: 20}}
              error = {numberrepairError ? true : false}
              id="numberrepair" 
              label="กรุณากรอกหมายเลข" 
              variant="standard"
              type="string"
              size="medium"
              helperText= {numberrepairError}
              value={numberrepair}
              onChange = {Numberrepairhandlechange}/>
              ):<FormControl disabled>
                  
              <InputLabel htmlFor="component-disabled"></InputLabel>
              <Input style={{  width: 150,marginRight: 20}} id="component-disabled" value={numberrepair} />
              <FormHelperText>ค้นหาสำเร็จ</FormHelperText>
              
    </FormControl>}
    <Button
                style={{ marginRight: 5}}
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SearchIcon/>}
                onClick={() => {
                  SearchAdminrepair();
                }}
              >
                ค้นหา
              </Button>
              <Button style={{ marginRight: 5}} variant="contained" size="large" onClick={() => {
                  cleardata();
                }} startIcon={<RestoreFromTrashIcon/>} >ลบ</Button>
              <Button style={{ marginRight: 5}} 
              variant="contained" color="secondary" size="large" href="/Selectadminrepair" startIcon={<ArrowBackIcon/>}> ย้อนกลับ </Button>
      </ContentHeader>
        <Container maxWidth="sm">
          <Grid container spacing={2}>
            <Grid item xs={12}></Grid>
          </Grid>
          
        </Container>
        {searchadminrepair ? (
        <TableContainer component={Paper}>
       <Table className={classes.table} aria-label="simple table">
         <TableHead>
           <TableRow>
             <TableCell align="center">หมายเลขบันทึกซ่อมของพนักงงาน</TableCell>
             <TableCell align="center">ชื่อ เจ้าหน้าที่</TableCell>
             <TableCell align="center">หมายเลขบันทึกการแจ้งซ่อม</TableCell>
             <TableCell align="center">ชื่อ อะไหล่ที่ใช้ซ่อมแซม</TableCell>
             <TableCell align="center">ความเสียหายที่เจ้าหน้าที่พบ</TableCell>
             <TableCell align="center">รายละเอียดการซ่อม</TableCell>
           </TableRow>
         </TableHead>
         <TableBody>
            {adminrepair.map((item: any) => (
             <TableRow key={item.id}>
               <TableCell align="center">{item.numberrepair}</TableCell>
               <TableCell align="center">{item.edges.AdminrepairPersonal.Personalname}</TableCell>
               <TableCell align="center">{item.edges.AdminrepairFix.Queue}</TableCell>
               <TableCell align="center">{item.edges.AdminrepairProduct.Productname}</TableCell>
               <TableCell align="center">{item.equipmentdamate}</TableCell>
               <TableCell align="center">{item.repairinformation}</TableCell>
             </TableRow>
           ))}
         </TableBody>
       </Table>
     </TableContainer>
    ): null}
    {searchall ? (
     <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">หมายเลขบันทึกซ่อมของพนักงงาน</TableCell>
           <TableCell align="center">ชื่อ เจ้าหน้าที่</TableCell>
           <TableCell align="center">หมายเลขบันทึกการแจ้งซ่อม</TableCell>
           <TableCell align="center">ชื่อ อะไหล่ที่ใช้ซ่อมแซม</TableCell>
           <TableCell align="center">ความเสียหายที่เจ้าหน้าที่พบ</TableCell>
           <TableCell align="center">รายละเอียดการซ่อม</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {adminrepairTable.map((item:any )=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.numberrepair}</TableCell>
             <TableCell align="center">{item.edges.adminrepairPersonal.personalname}</TableCell>
             <TableCell align="center">{item.edges.adminrepairFix.queue}</TableCell>
             <TableCell align="center">{item.edges.adminrepairProduct.productname}</TableCell>
             <TableCell align="center">{item.equipmentdamate}</TableCell>
             <TableCell align="center">{item.repairinformation}</TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   ):null}
      </Content>
    </Page>
  );
};