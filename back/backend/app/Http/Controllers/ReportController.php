<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use app\Models\Report;


class ReportController extends Controller
{
    public function index()
    {
        $reports = Report::all();
        return response()->json($reports, 200);
    }

    public function store(Request $request)
    {
        $request->validate([
            'user_id' => 'required|exists:users,id',
            'appointment_id' => 'required|exists:appointments,id',
            'procedure_id' => 'required|exists:procedures,id',
            'price_id' => 'required|exists:prices,id',
        ]);

        $report = Report::create($request->all());

        return response()->json($report, 201);
    }

    public function show(Report $report)
    {
        return response()->json($report, 200);
    }

    public function update(Request $request, Report $report)
    {
        $request->validate([
            'user_id' => 'required|exists:users,id',
            'appointment_id' => 'required|exists:appointments,id',
            'procedure_id' => 'required|exists:procedures,id',
            'price_id' => 'required|exists:prices,id',
        ]);

        $report->update($request->all());

        return response()->json($report, 200);
    }

    public function destroy(Report $report)
    {
        $report->delete();

        return response()->json(null, 204);
    }
}
