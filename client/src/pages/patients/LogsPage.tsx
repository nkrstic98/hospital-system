import {Log} from "../../types/Admission.ts";
import {ChangeEvent, useState} from "react";
import Paper from "@mui/material/Paper";
import TableContainer from "@mui/material/TableContainer";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import {TablePagination} from "@mui/material";
import {getTimeString} from "../../utils/utils.ts";
import Typography from "@mui/material/Typography";

export interface LogsPageProps {
    permission: string;
    logs: Log[]
}

export default function LogsPage({permission, logs}: LogsPageProps) {
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(10);

    const handleChangePage = (_: unknown, newPage: number) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event: ChangeEvent<HTMLInputElement>) => {
        setRowsPerPage(+event.target.value);
        setPage(0);
    };

    if (permission !== "READ") {
        return (
            <Typography variant="h4" gutterBottom component="div" sx={{mb: 3}}>
                You do not have permission to view this page
            </Typography>
        );
    }

    return (
        <Paper sx={{ width: '100%' }}>
            <TableContainer sx={{ maxHeight: 440 }}>
                <Table stickyHeader aria-label="sticky table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Timestamp</TableCell>
                            <TableCell>Action</TableCell>
                            <TableCell>Message</TableCell>
                            <TableCell>Performed By</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {logs
                            .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                            .map((row, index) => {
                                return (
                                    <>
                                        <TableRow hover tabIndex={-1} key={index}>
                                            <TableCell>{getTimeString(row.timestamp)}</TableCell>
                                            <TableCell>{row.action}</TableCell>
                                            <TableCell>{row.message}</TableCell>
                                            <TableCell>{row.performedBy}</TableCell>
                                        </TableRow>
                                    </>
                                );
                            })}
                    </TableBody>
                </Table>
            </TableContainer>
            <TablePagination
                rowsPerPageOptions={[10, 25, 100]}
                component="div"
                count={logs.length}
                rowsPerPage={rowsPerPage}
                page={page}
                onPageChange={handleChangePage}
                onRowsPerPageChange={handleChangeRowsPerPage}
            />
        </Paper>
    );
}