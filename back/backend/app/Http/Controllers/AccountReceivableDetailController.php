<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use app\Models\AccountReceivableDetail;

class AccountReceivableDetailController extends Controller
{
    public function index()
    {
        $details = AccountReceivableDetail::all();
        return response()->json($details, 200);
    }

    public function store(Request $request)
    {
        $request->validate([
            'user_id' => 'required|exists:users,id',
            'account_receivable_id' => 'required|exists:account_receivable,id',
        ]);

        $detail = AccountReceivableDetail::create($request->all());

        return response()->json($detail, 201);
    }

    public function show(AccountReceivableDetail $detail)
    {
        return response()->json($detail, 200);
    }

    public function update(Request $request, AccountReceivableDetail $detail)
    {
        $request->validate([
            'user_id' => 'required|exists:users,id',
            'account_receivable_id' => 'required|exists:account_receivable,id',
        ]);

        $detail->update($request->all());

        return response()->json($detail, 200);
    }

    public function destroy(AccountReceivableDetail $detail)
    {
        $detail->delete();

        return response()->json(null, 204);
    }
}
