<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class ProcedureController extends Controller
{
    public function index()
    {
        $procedures = Procedure::all();
        return response()->json($procedures, 200);
    }

    public function store(Request $request)
    {
        $request->validate([
            'name' => 'required|string|max:255',
            'description' => 'required|string|max:255',
        ]);

        $procedure = Procedure::create($request->all());

        return response()->json($procedure, 201);
    }

    public function show(Procedure $procedure)
    {
        return response()->json($procedure, 200);
    }

    public function update(Request $request, Procedure $procedure)
    {
        $request->validate([
            'name' => 'required|string|max:255',
            'description' => 'required|string|max:255',
        ]);

        $procedure->update($request->all());

        return response()->json($procedure, 200);
    }

    public function destroy(Procedure $procedure)
    {
        $procedure->delete();

        return response()->json(null, 204);
    }
}
