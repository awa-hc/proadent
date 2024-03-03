<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use app\Models\User;
use app\Models\Procedure;
use app\Models\Appointment;

class Price extends Model
{
    protected $fillable = ['user_id', 'procedure_id', 'appointment_days', 'total_price'];

    public function user()
    {
        return $this->belongsTo(User::class);
    }

    public function appointment()
    {
        return $this->belongsTo(Appointment::class);
    }

    public function procedure()
    {
        return $this->belongsTo(Procedure::class);
    }
}
