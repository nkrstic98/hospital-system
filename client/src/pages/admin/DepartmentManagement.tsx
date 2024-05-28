import Box from "@mui/material/Box";

const DepartmentManagement = () => {
    const daysOfWeek = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'];
    const shifts = ['Morning', 'Day', 'Night'];

    return (
        <Box sx={{display: 'flex', mt: 4}}>
            <div>
                <h2>Work Schedule</h2>
                <table>
                    <thead>
                    <tr>
                        <th>Day</th>
                        <th>Morning</th>
                        <th>Day</th>
                        <th>Night</th>
                    </tr>
                    </thead>
                    <tbody>
                    {daysOfWeek.map(day => (
                        <tr key={day}>
                            <td>{day}</td>
                            {shifts.map(shift => (
                                <td key={shift}>{shift}</td>
                            ))}
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>
        </Box>
    );
}

export default DepartmentManagement;