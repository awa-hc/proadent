<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class AccountReceivable extends Model
{
    protected $fillable = ['user_id', 'price_id', 'balance'];

    public function user()
    {
        return $this->belongsTo(User::class);
    }

    public function price()
    {
        return $this->belongsTo(Price::class);
    }
}
