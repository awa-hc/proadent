<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class AccountReceivableController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $procedures = Procedure::all();
        return view('procedures.index', compact('procedures'));
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        return view('procedures.create');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
    }

    /**
     * Display the specified resource.
     */
    public function show($id)
    {
        $procedure = Procedure::findOrFail($id);
        return view('procedures.show', compact('procedure'));
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit($id)
    {
        $procedure = Procedure::findOrFail($id);
        return view('procedures.edit', compact('procedure'));
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
    }
}
