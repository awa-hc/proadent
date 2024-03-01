<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use app\Models\AccountReceivable;

class AccountReceivableController extends Controller
{
    public function index()
    {
        $acounts = AccountReceivable::all();
        return response()->json($acounts, 200);
    }

    public function store(Request $request)
    {
        $request->validate([
            'user_id' => 'required|exists:users,id',
            'price_id' => 'required|exists:prices,id',
        ]);

        $acount = AccountReceivable::create($request->all());

        return response()->json($acount, 201);
    }

    public function show(AccountReceivable $acount)
    {
        return response()->json($acount, 200);
    }

    public function update(Request $request, AccountReceivable $acount)
    {
        $request->validate([
            'user_id' => 'required|exists:users,id',
            'price_id' => 'required|exists:prices,id',
        ]);

        $acount->update($request->all());

        return response()->json($acount, 200);
    }

    public function destroy(AccountReceivable $acount)
    {
        $acount->delete();

        return response()->json(null, 204);
    }
}
