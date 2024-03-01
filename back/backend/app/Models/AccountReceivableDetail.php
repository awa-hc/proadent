<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class AccountReceivableDetail extends Model
{
    protected $fillable = ['user_id', 'account_receivable_id', 'amount'];

    public function user()
    {
        return $this->belongsTo(User::class);
    }

    public function accountReceivable()
    {
        return $this->belongsTo(AccountReceivable::class);
    }
}
