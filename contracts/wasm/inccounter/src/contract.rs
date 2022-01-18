// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use std::ptr;

use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;

pub struct CallIncrementCall {
	pub func: ScFunc,
}

pub struct CallIncrementRecurse5xCall {
	pub func: ScFunc,
}

pub struct EndlessLoopCall {
	pub func: ScFunc,
}

pub struct IncrementCall {
	pub func: ScFunc,
}

pub struct IncrementWithDelayCall {
	pub func: ScFunc,
	pub params: MutableIncrementWithDelayParams,
}

pub struct InitCall {
	pub func: ScInitFunc,
	pub params: MutableInitParams,
}

pub struct LocalStateInternalCallCall {
	pub func: ScFunc,
}

pub struct LocalStatePostCall {
	pub func: ScFunc,
}

pub struct LocalStateSandboxCallCall {
	pub func: ScFunc,
}

pub struct PostIncrementCall {
	pub func: ScFunc,
}

pub struct RepeatManyCall {
	pub func: ScFunc,
	pub params: MutableRepeatManyParams,
}

pub struct TestVliCodecCall {
	pub func: ScFunc,
}

pub struct TestVluCodecCall {
	pub func: ScFunc,
}

pub struct WhenMustIncrementCall {
	pub func: ScFunc,
	pub params: MutableWhenMustIncrementParams,
}

pub struct GetCounterCall {
	pub func: ScView,
	pub results: ImmutableGetCounterResults,
}

pub struct GetVliCall {
	pub func: ScView,
	pub params: MutableGetVliParams,
	pub results: ImmutableGetVliResults,
}

pub struct GetVluCall {
	pub func: ScView,
	pub params: MutableGetVluParams,
	pub results: ImmutableGetVluResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn call_increment(_ctx: & dyn ScFuncCallContext) -> CallIncrementCall {
        CallIncrementCall {
            func: ScFunc::new(HSC_NAME, HFUNC_CALL_INCREMENT),
        }
    }

    pub fn call_increment_recurse5x(_ctx: & dyn ScFuncCallContext) -> CallIncrementRecurse5xCall {
        CallIncrementRecurse5xCall {
            func: ScFunc::new(HSC_NAME, HFUNC_CALL_INCREMENT_RECURSE5X),
        }
    }

    pub fn endless_loop(_ctx: & dyn ScFuncCallContext) -> EndlessLoopCall {
        EndlessLoopCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ENDLESS_LOOP),
        }
    }

    pub fn increment(_ctx: & dyn ScFuncCallContext) -> IncrementCall {
        IncrementCall {
            func: ScFunc::new(HSC_NAME, HFUNC_INCREMENT),
        }
    }

    pub fn increment_with_delay(_ctx: & dyn ScFuncCallContext) -> IncrementWithDelayCall {
        let mut f = IncrementWithDelayCall {
            func: ScFunc::new(HSC_NAME, HFUNC_INCREMENT_WITH_DELAY),
            params: MutableIncrementWithDelayParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }

    pub fn init(_ctx: & dyn ScFuncCallContext) -> InitCall {
        let mut f = InitCall {
            func: ScInitFunc::new(HSC_NAME, HFUNC_INIT),
            params: MutableInitParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }

    pub fn local_state_internal_call(_ctx: & dyn ScFuncCallContext) -> LocalStateInternalCallCall {
        LocalStateInternalCallCall {
            func: ScFunc::new(HSC_NAME, HFUNC_LOCAL_STATE_INTERNAL_CALL),
        }
    }

    pub fn local_state_post(_ctx: & dyn ScFuncCallContext) -> LocalStatePostCall {
        LocalStatePostCall {
            func: ScFunc::new(HSC_NAME, HFUNC_LOCAL_STATE_POST),
        }
    }

    pub fn local_state_sandbox_call(_ctx: & dyn ScFuncCallContext) -> LocalStateSandboxCallCall {
        LocalStateSandboxCallCall {
            func: ScFunc::new(HSC_NAME, HFUNC_LOCAL_STATE_SANDBOX_CALL),
        }
    }

    pub fn post_increment(_ctx: & dyn ScFuncCallContext) -> PostIncrementCall {
        PostIncrementCall {
            func: ScFunc::new(HSC_NAME, HFUNC_POST_INCREMENT),
        }
    }

    pub fn repeat_many(_ctx: & dyn ScFuncCallContext) -> RepeatManyCall {
        let mut f = RepeatManyCall {
            func: ScFunc::new(HSC_NAME, HFUNC_REPEAT_MANY),
            params: MutableRepeatManyParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }

    pub fn test_vli_codec(_ctx: & dyn ScFuncCallContext) -> TestVliCodecCall {
        TestVliCodecCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TEST_VLI_CODEC),
        }
    }

    pub fn test_vlu_codec(_ctx: & dyn ScFuncCallContext) -> TestVluCodecCall {
        TestVluCodecCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TEST_VLU_CODEC),
        }
    }

    pub fn when_must_increment(_ctx: & dyn ScFuncCallContext) -> WhenMustIncrementCall {
        let mut f = WhenMustIncrementCall {
            func: ScFunc::new(HSC_NAME, HFUNC_WHEN_MUST_INCREMENT),
            params: MutableWhenMustIncrementParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }

    pub fn get_counter(_ctx: & dyn ScViewCallContext) -> GetCounterCall {
        let mut f = GetCounterCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_COUNTER),
            results: ImmutableGetCounterResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }

    pub fn get_vli(_ctx: & dyn ScViewCallContext) -> GetVliCall {
        let mut f = GetVliCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_VLI),
            params: MutableGetVliParams { id: 0 },
            results: ImmutableGetVliResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }

    pub fn get_vlu(_ctx: & dyn ScViewCallContext) -> GetVluCall {
        let mut f = GetVluCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_VLU),
            params: MutableGetVluParams { id: 0 },
            results: ImmutableGetVluResults { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, &mut f.results.id);
        f
    }
}
