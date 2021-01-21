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
import { DefaultApi } from '../../api/apis';
import { EntAdminrepair } from '../../api/models/EntAdminrepair';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [loading, setLoading] = useState(true);
 const [adminrepairs, setAdminrepairs] = useState<EntAdminrepair[]>([]);
 
 useEffect(() => {
    const getAdminrepairs = async () => {
      const res:any = await api.listAdminrepair({ limit: 10, offset: 0 });
      setLoading(true);
      setAdminrepairs(res);
      console.log(res);
    };
    getAdminrepairs();
  }, [loading]);
 
 const deleteAdminrepairs = async (id: number) => {
   const res = await api.deleteAdminrepair({ id: id });
   setLoading(true);
   window.location.reload(false);
 };
 
 return (
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
         {adminrepairs.map((item:any )=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.numberrepair}</TableCell>
             <TableCell align="center">{item.edges.adminrepairPersonal.personalname}</TableCell>
             <TableCell align="center">{item.edges.adminrepairFix.queue}</TableCell>
             <TableCell align="center">{item.edges.adminrepairProduct.productname}</TableCell>
             <TableCell align="center">{item.equipmentdamate}</TableCell>
             <TableCell align="center">{item.repairinformation}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteAdminrepairs(item.id);
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
 );
}
