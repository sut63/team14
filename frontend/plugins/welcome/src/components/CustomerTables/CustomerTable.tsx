import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import moment from 'moment';
import { EntCustomer } from '../../api/models/EntCustomer';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import FaceIcon from '@material-ui/icons/Face';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function CustomerTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [customers, setCustomers] = useState<EntCustomer[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getCustomers = async () => {
     const res = await api.listCustomer({ limit: 10, offset: 0 });
     setLoading(false);
     setCustomers(res);
     console.log(res);
   };
   getCustomers();
 }, [loading]);

const deleteCustomer = async (id: number) => {
 const res = await api.deleteCustomer({ id: id });
 setLoading(true);
};

const getCustomer = async (id: number) => {
  const res = await api.getCustomer({ id: id });
  setLoading(true);
 };
 
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`Customer System`} type="Computer Repair System" >
      <div>&nbsp;&nbsp;&nbsp;</div>
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

    <ContentHeader title="ตารางแสดงผลข้อมูล">
      <div>&nbsp;&nbsp;&nbsp;</div>
      <Button variant="contained" color="secondary" href="/CustomerSearch" startIcon={<FaceIcon />}> ค้นหาข้อมูลลูกค้า </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="secondary" href="/CreateNewCustomer" startIcon={<FaceIcon />}> เพิ่มข้อมูลลูกค้า </Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/Group14" startIcon={<HomeRoundedIcon/>}> กลับหน้าหลัก </Button>
      </ContentHeader>

   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ลำดับที่</TableCell>
           <TableCell align="center">คำนำหน้าชื่อ</TableCell>
           <TableCell align="center">ชื่อ-นามสกุล</TableCell>
           <TableCell align="center">ที่อยู่</TableCell>
           <TableCell align="center">เบอร์โทร</TableCell>
           <TableCell align="center">เลขบัตรประจำตัวประชาชน</TableCell>
           <TableCell align="center">เพศ</TableCell>
           <TableCell align="center">เจ้าหน้าที่แจ้งซ่อม</TableCell>

         </TableRow>
       </TableHead>
       <TableBody>
         {customers.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.title?.titlename}</TableCell>
             <TableCell align="center">{item.customername}</TableCell>
             <TableCell align="center">{item.address}</TableCell>
             <TableCell align="center">{item.phonenumber}</TableCell>
             <TableCell align="center">{item.identificationnumber}</TableCell>
             <TableCell align="center">{item.edges?.gender?.gendername}</TableCell>
             <TableCell align="center">{item.edges?.personal.personalname}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                    deleteCustomer(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 DELETE
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
 
